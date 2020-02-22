package notification

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

// GetHistory ...
func (gateway *NotificationGateway) GetHistory(msisdn string) ([]entity.DynamoItemHistory, error) {
	data, err := gateway.AwsLib.GetDynamoHistory(msisdn)
	dynamoItem := []entity.DynamoItemHistory{}
	err = dynamodbattribute.UnmarshalListOfMaps(data.Items, &dynamoItem)
	return dynamoItem, err
}

// GetByID ...
func (gateway *NotificationGateway) GetByID(ID string) (*entity.DynamoItemHistory, error) {
	data, err := gateway.AwsLib.GetDynamoData(ID)
	dynamoItem := &entity.DynamoItemHistory{}
	err = dynamodbattribute.UnmarshalMap(data.Item, &dynamoItem)
	if err != nil {
		return nil, err
	}
	return dynamoItem, err
}
