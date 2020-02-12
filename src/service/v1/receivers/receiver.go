package receivers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
	"github.com/sofyan48/rll-daemon-new/src/service/v1/receivers/callbackprovider"
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
	InfobipReceiver(ID string, data *entity.InfobipCallbackRequest) (string, error)
	WavecellReceiver(ID string, data *entity.WavecelllCallBackRequest) (string, error)
}

// InfobipReceiver ...
func (rcv *Receiver) InfobipReceiver(ID string, data *entity.InfobipCallbackRequest) (string, error) {
	dynamoItem := &entity.DynamoItemResponse{}
	fmt.Println(data)
	dynamoData, err := rcv.AwsLib.GetDynamoData(ID)
	if err != nil {
		return "", err
	}
	err = dynamodbattribute.UnmarshalMap(dynamoData.Item, &dynamoItem)
	if err != nil {
		return "", err
	}
	rcv.Callback.InfobipCallback(dynamoItem, data)
	return "", nil
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
