package transmiter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/service/kinesis"
	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/libaws"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/provider"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/request"
)

// Transmiter ...
type Transmiter struct {
	AwsLibs   libaws.AwsInterface
	Provider  provider.ProvidersInterface
	Requester request.RequesterInterface
}

// GetTransmiter ...
func GetTransmiter() *Transmiter {
	return &Transmiter{
		AwsLibs:   libaws.AwsHAndler(),
		Provider:  provider.ProvidersHandler(),
		Requester: request.RequesterHandler(),
	}
}

// ConsumerTrans ...
func (trs *Transmiter) ConsumerTrans(wg *sync.WaitGroup) {
	shardIterator, err := trs.AwsLibs.GetShardIterator()
	if err != nil {
		log.Println(err)
	}

	describeInput := trs.AwsLibs.GetDescribeInput()
	describeInput.SetStreamName("notification")
	describeInput.SetExclusiveStartShardId(os.Getenv("KINESIS_SHARD_ID"))
	for {
		err := trs.AwsLibs.WaitUntil(describeInput)
		if err != nil {
			log.Println("error Wait: ", err)
		}
		done := make(chan bool)
		go func() {
			msgInput := &kinesis.GetRecordsInput{}
			msgInput.SetShardIterator(shardIterator)

			data, err := trs.AwsLibs.Consumer(msgInput)
			if err != nil {
				log.Println(err)
			}
			itemDynamo := &entity.DynamoItem{}
			for _, i := range data.Records {
				err := json.Unmarshal([]byte(string(i.Data)), itemDynamo)
				if err != nil {
					log.Println("Error: ", err)
				}
				trs.intercepActionShard(itemDynamo)
			}
			close(done)
			shardIterator = *data.NextShardIterator
			return
		}()
		<-done
		time.Sleep(5 * time.Second)
	}

}

func (trs *Transmiter) intercepActionShard(data *entity.DynamoItem) {
	dataThirdParty := make([]entity.DataProvider, 0)
	err := json.Unmarshal([]byte(os.Getenv("SMS_ORDER_CONF")), &dataThirdParty)
	if err != nil {
		log.Println(err)
	}
	history := &entity.HistoryItem{}
	var historyProvider string
	for _, i := range dataThirdParty {
		if data.History[i.Provider] != nil {
			history.CallbackData = data.History[i.Provider].CallbackData
			history.DeliveryReport = data.History[i.Provider].DeliveryReport
			history.Payload = data.History[i.Provider].Payload
			history.Response = data.History[i.Provider].Response
			historyProvider = data.History[i.Provider].Provider
			break
		}
	}
	trs.wavecellActionShard(historyProvider, history)

}

func (trs *Transmiter) infobipActionShardOTP(history string, payload *entity.HistoryItem) {
	dest := entity.InfobipDestination{}
	dest.To = payload.Payload.Msisdn
	destination := []entity.InfobipDestination{dest}

	infobipMessages := entity.InfobipMessages{}
	infobipMessages.Destinations = destination
	infobipMessages.From = os.Getenv("INFOBIP_SENDER_ID")
	infobipMessages.Text = payload.Payload.Text
	infobipMessages.NotifyContenType = "application/json"
	infobipMessages.NotifyURL = os.Getenv("INFOBIP_CALLBACK")
	infobipMessages.CallbackData = payload.CallbackData
	infobipMessagesSlice := []entity.InfobipMessages{infobipMessages}

	infobip := &entity.InfobipRequestPayload{}
	infobip.Messages = infobipMessagesSlice
	reformatPayload, err := json.Marshal(infobip)
	if err != nil {
		log.Println("Error: ", err)
	}
	username := os.Getenv("INFOBIP_USERNAME")
	password := os.Getenv("INFOBIP_PASSWORD")
	client, err := trs.Requester.CLIENT("POST", os.Getenv("INFOBIP_SEND_SMS_URL"), reformatPayload)
	if err != nil {
		log.Println("Error: ", err)
	}
	requester := &http.Client{}
	client.SetBasicAuth(username, password)
	client.Header.Set("Content-Type", "application/json")
	response, err := requester.Do(client)
	if err != nil {
		log.Println("Infobip Transmitter: ", err)
	}
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
	body, err := ioutil.ReadAll(response.Body)
	s := string(body)
	fmt.Println(s)
}

func (trs *Transmiter) wavecellActionShard(history string, payload *entity.HistoryItem) {
	reformatPayload := &entity.WavecellRequest{}
	reformatPayload.Destination = payload.Payload.Msisdn
	reformatPayload.Source = os.Getenv("WAVECELL_ACC_ID")
	reformatPayload.Text = payload.Payload.Text
	reformatPayload.DLRCallback = os.Getenv("WAVECELL_CALLBACK_URL")
	wavecelSendURL := "https://api.wavecell.com/sms/v1/" + os.Getenv("WAVECELL_SUB_ACC_ID_GENERAL") + "/single"
	if payload.Payload.OTP == true {
		wavecelSendURL = "https://api.wavecell.com/sms/v1/" + os.Getenv("WAVECELL_SUB_ACC_ID") + "/single"
	}
	wavecellReformatPayload, err := json.Marshal(reformatPayload)
	client, err := trs.Requester.CLIENT("POST", wavecelSendURL, wavecellReformatPayload)
	if err != nil {
		log.Println("Error: ", err)
	}
	requester := &http.Client{}
	client.Header.Set("Content-Type", "application/json")
	client.Header.Set("Authorization", "Bearer "+os.Getenv("WAVECELL_ACC_TOKEN"))
	response, err := requester.Do(client)
	if err != nil {
		log.Println("Wavecell Transmitter: ", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Wavecell Transmitter: ", err)
	}
	wavecellResponse := &entity.WavecellResponse{}
	json.Unmarshal(body, wavecellResponse)
	_, err = trs.updateDynamoTransmitt(payload.CallbackData,
		wavecellResponse.Status.Code,
		string(body))
	if err != nil {
		log.Println("Wavecell Transmitter Dynamo: ", err)
	}
}

// updateDynamoTransmitt ...
func (trs *Transmiter) updateDynamoTransmitt(ID, status, data string) (string, error) {
	result, err := trs.AwsLibs.UpdateDynamo(ID, status, data)
	return result.GoString(), err
}

// TransferToShardReceiver ...
func (trs *Transmiter) TransferToShardReceiver(historyString string) {}
