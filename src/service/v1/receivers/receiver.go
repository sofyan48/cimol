package receivers

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/libaws"
)

// Receiver ...
type Receiver struct {
	AwsLib libaws.AwsInterface
}

// ReceiverHandler ...
func ReceiverHandler() *Receiver {
	return &Receiver{
		AwsLib: libaws.AwsHAndler(),
	}
}

// ReceiverInterface ...
type ReceiverInterface interface {
	InfobipReceiver(ID string, data *entity.InfobipCallBackRequest) (string, error)
	WavecellReceiver(ID string, data *entity.WavecelllCallBackRequest) (string, error)
}

// InfobipReceiver ...
func (rcv *Receiver) InfobipReceiver(ID string, data *entity.InfobipCallBackRequest) (string, error) {
	dynamoItem := &entity.DynamoItemResponse{}
	dynamoData, err := rcv.AwsLib.GetDynamoData(ID)
	if err != nil {
		return "", err
	}
	err = dynamodbattribute.UnmarshalMap(dynamoData.Item, &dynamoItem)
	if err != nil {
		return "", err
	}
	return "", nil
}

// WavecellReceiver ...
func (rcv *Receiver) WavecellReceiver(ID string, data *entity.WavecelllCallBackRequest) (string, error) {
	fmt.Println(data)
	return "", nil
}
