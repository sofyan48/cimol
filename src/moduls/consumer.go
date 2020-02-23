package moduls

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/service/kinesis"
	entity "github.com/sofyan48/cimol/src/entity/http/v1"
	"github.com/sofyan48/cimol/src/moduls/notification/email"
	"github.com/sofyan48/cimol/src/moduls/notification/sms"
	"github.com/sofyan48/cimol/src/util/helper/libaws"
	"github.com/sofyan48/cimol/src/util/logging"
)

// ModulsConsumer ...
type ModulsConsumer struct {
	AwsLibs libaws.AwsInterface
	Logs    logging.LogInterface
	SMS     sms.SMSModulsInterface
	Email   email.EmailTransmiterInterface
}

// GetModuls ...
func GetModuls() *ModulsConsumer {
	return &ModulsConsumer{
		AwsLibs: libaws.AwsHAndler(),
		Logs:    logging.LogHandler(),
		SMS:     sms.SMSModulsHandler(),
		Email:   email.EmailModulsHandler(),
	}
}

// MainModuls ...
func (modul *ModulsConsumer) MainModuls(wg *sync.WaitGroup) {
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
				switch selectionType.Type {
				case "sms":
					itemSMS := &entity.DynamoItem{}
					json.Unmarshal([]byte(string(i.Data)), itemSMS)
					modul.SMS.IntercepActionShardSMS(itemSMS)
				case "email":
					itemEmail := &entity.DynamoItemEmail{}
					json.Unmarshal([]byte(string(i.Data)), itemEmail)
					modul.Email.IntercepActionShardEmail(itemEmail)
				}
			}
			close(done)
			shardIterator = *data.NextShardIterator
			return
		}()
		<-done
		time.Sleep(3 * time.Second)
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
