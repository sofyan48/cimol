package receivers

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	entity "github.com/sofyan48/otp/src/entity/http/v1"
	"github.com/sofyan48/otp/src/util/callbackprovider"
	"github.com/sofyan48/otp/src/util/helper/libaws"
)

// Receiver ...
type Receiver struct {
	AwsLib   libaws.AwsInterface
	Callback callbackprovider.ProviderCallbackInterface
}

// ReceiverHandler ...
func ReceiverHandler() *Receiver {
	return &Receiver{
		AwsLib:   libaws.AwsHAndler(),
		Callback: callbackprovider.ProviderCallbackHandler(),
	}
}

// ReceiverInterface ...
type ReceiverInterface interface {
	InfobipReceiver(ID string, data *entity.InfobipCallbackRequest)
	WavecellReceiver(ID string, data *entity.WavecelllCallBackRequest) (string, error)
}

// InfobipReceiver ...
func (rcv *Receiver) InfobipReceiver(ID string, data *entity.InfobipCallbackRequest) {
	dynamoItem := &entity.DynamoItemResponse{}
	dynamoData, err := rcv.AwsLib.GetDynamoData(ID)
	if err != nil {
		log.Println("Error: ", err)
	}
	err = dynamodbattribute.UnmarshalMap(dynamoData.Item, &dynamoItem)
	if err != nil {
		log.Println("Error: ", err)
	}
	historyItems := &entity.HistoryItem{}
	err = json.Unmarshal([]byte(dynamoItem.History[1]), historyItems)
	if err != nil {
		log.Println("Error: ", err)
	}

	rcv.Callback.InfobipCallback(dynamoItem, data, historyItems)
}

// WavecellReceiver ...
func (rcv *Receiver) WavecellReceiver(ID string, data *entity.WavecelllCallBackRequest) (string, error) {

	return "", nil
}

// GoSMSReceiver ...
func (rcv *Receiver) GoSMSReceiver(ID string, data *entity.WavecelllCallBackRequest) (string, error) {

	return "", nil
}
