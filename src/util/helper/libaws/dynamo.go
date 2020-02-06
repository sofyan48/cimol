package libaws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	dynamoEntyty "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
)

// GetDynamoDB get dynamodb service
// return *dynamodb.DynamoDB
func (aw *Aws) GetDynamoDB() *dynamodb.DynamoDB {
	dynamo := dynamodb.New(session.New(), aw.Config)
	return dynamo
}

// InputDynamo ...
func (aw *Aws) InputDynamo(itemDynamo *dynamoEntyty.DynamoItem) (*dynamodb.PutItemOutput, error) {
	dynamoLibs := aw.GetDynamoDB()
	result := &dynamodb.PutItemOutput{}
	mItem, err := dynamodbattribute.MarshalMap(itemDynamo)
	if err != nil {
		return result, err
	}
	inputDynamo := &dynamodb.PutItemInput{
		Item:      mItem,
		TableName: aws.String(os.Getenv("AWS_DYNAMO_TABLE")),
	}
	result, err = dynamoLibs.PutItem(inputDynamo)
	if err != nil {
		return result, err
	}
	return result, nil
}

// UpdateDynamo ...
func (aw *Aws) UpdateDynamo(ID string, itemDynamo *dynamoEntyty.DynamoItem) (*dynamodb.UpdateItemOutput, error) {
	dynamoLibs := aw.GetDynamoDB()

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(ID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":valhistory": {
				N: aws.String(itemDynamo.History),
			},
		},
		TableName:        aws.String(os.Getenv("AWS_DYNAMO_TABLE")),
		ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String("SET #hsty = :valhistory, #sts = :valstatusText, updatedAt = :updatedAt"),
	}
	result, err := dynamoLibs.UpdateItem(input)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetDynamoData ..
func (aw *Aws) GetDynamoData(ID string) (*dynamoEntyty.DynamoItem, error) {
	dynamoLibs := aw.GetDynamoDB()
	result, err := dynamoLibs.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("AWS_DYNAMO_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(ID),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	itemDynamo := &dynamoEntyty.DynamoItem{}
	err = dynamodbattribute.UnmarshalMap(result.Item, itemDynamo)
	if err != nil {
		return nil, err
	}
	return itemDynamo, nil
}

// GetDynamoHistory ..
func (aw *Aws) GetDynamoHistory(receiverAddress string) (*dynamoEntyty.DynamoItem, error) {
	dynamoLibs := aw.GetDynamoDB()
	result, err := dynamoLibs.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("AWS_DYNAMO_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"receiverAddress": {
				N: aws.String(receiverAddress),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	itemDynamo := &dynamoEntyty.DynamoItem{}
	err = dynamodbattribute.UnmarshalMap(result.Item, itemDynamo)
	if err != nil {
		return nil, err
	}
	return itemDynamo, nil
}
