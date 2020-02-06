package gateway

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/libaws"
)

// Gateway ...
type Gateway struct {
	AwsLib libaws.AwsInterface
}

// GatewayHandler Handler
func GatewayHandler() *Gateway {
	return &Gateway{
		AwsLib: libaws.AwsHAndler(),
	}
}

//GatewayInterface declare All Method
type GatewayInterface interface {
	PostNotification() *entity.PostNotificationResponse
	GetHistory(msisdn string) (string, error)
	GetByID(ID string) (string, error)
}

// PostNotification ...
// return *entity.PostNotificationResponse
func (gateway *Gateway) PostNotification() *entity.PostNotificationResponse {
	result := &entity.PostNotificationResponse{}
	result.ID = "ID"
	result.Status = "Status"
	return result
}

// GetHistory ...
func (gateway *Gateway) GetHistory(msisdn string) (string, error) {
	data, err := gateway.AwsLib.GetDynamoHistory(msisdn)
	return data.GoString(), err
}

// GetByID ...
func (gateway *Gateway) GetByID(ID string) (*entity.DynamoItem, error) {
	data, err := gateway.AwsLib.GetDynamoData(ID)
	dynamoItem := entity.Testing{}
	fmt.Println(data)
	err = dynamodbattribute.UnmarshalMap(data, &dynamoItem)
	if err != nil {
		return nil, err
	}

	return nil, err
}
