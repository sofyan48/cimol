package gateway

import (
	"sync"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/libaws"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/request"
)

// Gateway ...
type Gateway struct {
	AwsLib    libaws.AwsInterface
	Requester request.RequesterInterface
}

// GatewayHandler Handler
func GatewayHandler() *Gateway {
	return &Gateway{
		AwsLib:    libaws.AwsHAndler(),
		Requester: request.RequesterHandler(),
	}
}

//GatewayInterface declare All Method
type GatewayInterface interface {
	PostNotification(data *entity.PostNotificationRequest, wg *sync.WaitGroup)
	GetHistory(msisdn string) ([]entity.DynamoItemResponse, error)
	GetByID(ID string) (*entity.DynamoItemResponse, error)
}

// PostNotification ...
// return *entity.PostNotificationResponse
func (gateway *Gateway) PostNotification(data *entity.PostNotificationRequest, wg *sync.WaitGroup) {
	itemDynamo := &entity.DynamoItem{}
	itemDynamo.Data = data.Payload.Text
	itemDynamo.ReceiverAddress = data.Payload.Msisdn
	itemDynamo.History = ""
	itemDynamo.StatusText = "QUEUE"
	itemDynamo.ID = data.UUID
	itemDynamo.Type = data.Type
	wg.Add(2)
	go gateway.AwsLib.InputDynamo(itemDynamo, wg)
	wg.Done()
}

// GetHistory ...
func (gateway *Gateway) GetHistory(msisdn string) ([]entity.DynamoItemResponse, error) {
	data, err := gateway.AwsLib.GetDynamoHistory(msisdn)
	dynamoItem := []entity.DynamoItemResponse{}
	err = dynamodbattribute.UnmarshalListOfMaps(data.Items, &dynamoItem)
	return dynamoItem, err
}

// GetByID ...
func (gateway *Gateway) GetByID(ID string) (*entity.DynamoItemResponse, error) {
	data, err := gateway.AwsLib.GetDynamoData(ID)
	dynamoItem := &entity.DynamoItemResponse{}
	err = dynamodbattribute.UnmarshalMap(data.Item, &dynamoItem)
	if err != nil {
		return nil, err
	}
	return dynamoItem, err
}
