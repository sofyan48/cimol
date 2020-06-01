package moduls

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/aws/aws-sdk-go/service/kinesis"
	entity "github.com/sofyan48/cimol/src/entity/http/v1"
	"github.com/sofyan48/cimol/src/moduls/notification/email"
	"github.com/sofyan48/cimol/src/moduls/notification/sms"
	"github.com/sofyan48/cimol/src/util/helper/libaws"
	"github.com/sofyan48/cimol/src/util/kafka"
	"github.com/sofyan48/cimol/src/util/logging"
)

// ModulsConsumer ...
type ModulsConsumer struct {
	AwsLibs libaws.AwsInterface
	Logs    logging.LogInterface
	SMS     sms.SMSModulsInterface
	Email   email.EmailTransmiterInterface
	Kafka   kafka.KafkaLibraryInterface
}

// GetModuls ...
func GetModuls() *ModulsConsumer {
	return &ModulsConsumer{
		AwsLibs: libaws.AwsHAndler(),
		Logs:    logging.LogHandler(),
		SMS:     sms.SMSModulsHandler(),
		Email:   email.EmailModulsHandler(),
		Kafka:   kafka.KafkaLibraryHandler(),
	}
}

// MainModuls ...
func (modul *ModulsConsumer) MainModuls(wg *sync.WaitGroup) {
	switch os.Getenv("BROKER_MODULS") {
	case "kinesis":
		modul.kinesisModuls()
	case "kafka":
		modul.kafkaModuls()
	}
}

// kafkaModuls ...
func (modul *ModulsConsumer) kafkaModuls(topics string) {
	client, err := modul.Kafka.InitNewConsumer()
	if err != nil {
		panic(err)
	}
	response, err := client.ConsumePartition(topics, 0, sarama.OffsetNewest)
	if err != nil {
		panic(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	doneCh := make(chan struct{})
	go func() {
		for {
			eventData := &entity.StateFullFormatKafka{}
			select {
			case err := <-response.Errors():
				fmt.Println(err)
			case message := <-response.Messages():
				log.Println("EV Receive: ", message.Timestamp, " | Topic: ", message.Topic)
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh
}

func (modul *ModulsConsumer) kinesisModuls() {
	shardIterator, err := modul.AwsLibs.GetShardIterator()
	if err != nil {
		modul.Logs.Write("Transmitter", err.Error())
	}

	describeInput := modul.AwsLibs.GetDescribeInput()
	describeInput.SetStreamName("notification")
	describeInput.SetExclusiveStartShardId(os.Getenv("KINESIS_SHARD_ID"))
	for {
		err := modul.AwsLibs.WaitUntil(describeInput)
		if err != nil {
			modul.Logs.Write("Transmitter", err.Error())
		}
		done := make(chan bool)
		go func() {
			msgInput := &kinesis.GetRecordsInput{}
			msgInput.SetShardIterator(shardIterator)
			data, err := modul.AwsLibs.Consumer(msgInput)
			if err != nil {
				modul.Logs.Write("Transmitter", err.Error())
			}
			selectionType := &entity.DataReceiveSelection{}
			for _, i := range data.Records {
				json.Unmarshal([]byte(string(i.Data)), selectionType)
				fmt.Println("Receive: ", selectionType.Type)
				modul.kinesisTransmitter(selectionType, i)
			}
			close(done)
			shardIterator = *data.NextShardIterator
			return
		}()
		<-done
		time.Sleep(3 * time.Second)
	}
}

func (modul *ModulsConsumer) kinesisTransmitter(selectionType *entity.DataReceiveSelection, data *kinesis.Record) {
	switch selectionType.Type {
	case "sms":
		itemSMS := &entity.DynamoItem{}
		json.Unmarshal([]byte(string(data.Data)), itemSMS)
		modul.SMS.IntercepActionShardSMS(itemSMS)
	case "email":
		itemEmail := &entity.DynamoItemEmail{}
		json.Unmarshal([]byte(string(data.Data)), itemEmail)
		modul.Email.IntercepActionShardEmail(itemEmail)
	}
}

// updateDynamoTransmitt ...
func (modul *ModulsConsumer) updateDynamoTransmitt(ID, status, data string, history *entity.HistoryItem) (string, error) {
	result, err := modul.AwsLibs.UpdateDynamo(ID, status, data, history)
	return result.GoString(), err
}

func checkEnvironment() bool {
	envi := os.Getenv("APP_ENVIRONMENT")
	return envi == "development" || envi == "staging"
}
