package gateway

import (
	"fmt"

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
	PostNotification(data *entity.PostNotificationRequest) (*entity.PostNotificationResponse, error)
	GetHistory(msisdn string) ([]entity.DynamoItemResponse, error)
	GetByID(ID string) (*entity.DynamoItemResponse, error)
}

// PostNotification ...
// return *entity.PostNotificationResponse
func (gateway *Gateway) PostNotification(data *entity.PostNotificationRequest) (*entity.PostNotificationResponse, error) {
	result := &entity.PostNotificationResponse{}
	fmt.Println(data.Payload.Msisdn)
	return result, nil
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
