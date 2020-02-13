package receivers

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
	"github.com/sofyan48/rll-daemon-new/src/util/callbackprovider"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/libaws"
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
	rcv.Callback.InfobipCallback(dynamoItem, data)
}

// WavecellReceiver ...
func (rcv *Receiver) WavecellReceiver(ID string, data *entity.WavecelllCallBackRequest) (string, error) {
	fmt.Println(data)
	return "", nil
}

// GoSMSReceiver ...
func (rcv *Receiver) GoSMSReceiver(ID string, data *entity.WavecelllCallBackRequest) (string, error) {
	fmt.Println(data)
	return "", nil
}
