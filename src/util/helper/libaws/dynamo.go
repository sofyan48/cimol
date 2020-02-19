package libaws

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	dynamoEntyty "github.com/sofyan48/cimol/src/entity/http/v1"
)

// GetDynamoDB get dynamodb service
// return *dynamodb.DynamoDB
func (aw *Aws) GetDynamoDB() *dynamodb.DynamoDB {
	cfg := aw.Sessions()
	dynamo := dynamodb.New(session.New(), cfg)
	return dynamo
}

// InputDynamo ...
func (aw *Aws) InputDynamo(itemDynamo *dynamoEntyty.DynamoItem, wg *sync.WaitGroup) (*dynamodb.PutItemOutput, error) {
	dynamoLibs := aw.GetDynamoDB()
	result := &dynamodb.PutItemOutput{}
	mItem, err := dynamodbattribute.MarshalMap(itemDynamo)
	if err != nil {
		log.Println("ERROR: ", err)
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
	wg.Done()
	return result, nil
}

// InputDynamoEmail ...
func (aw *Aws) InputDynamoEmail(itemDynamo *dynamoEntyty.DynamoItemEmail, wg *sync.WaitGroup) (*dynamodb.PutItemOutput, error) {
	dynamoLibs := aw.GetDynamoDB()
	result := &dynamodb.PutItemOutput{}
	mItem, err := dynamodbattribute.MarshalMap(itemDynamo)
	if err != nil {
		log.Println("ERROR: ", err)
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
	wg.Done()
	return result, nil
}

// UpdateDynamo ...
func (aw *Aws) UpdateDynamo(ID, status, data string, history *dynamoEntyty.HistoryItem) (*dynamodb.UpdateItemOutput, error) {
	dynamoLibs := aw.GetDynamoDB()
	historyReformat, err := json.Marshal(history)
	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(ID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":valhistory": {
				SS: aws.StringSlice([]string{
					string(historyReformat),
					data,
				}),
			},
			":valstatusText": {
				S: aws.String(status),
			},
		},

		TableName:        aws.String(os.Getenv("AWS_DYNAMO_TABLE")),
		ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String("SET history = :valhistory, statusText = :valstatusText"),
	}
	result, err := dynamoLibs.UpdateItem(input)
	if err != nil {
		return result, err
	}
	return result, nil
}

// CallbackSendUpdate ...
func (aw *Aws) CallbackSendUpdate(ID, status, oldHistory string, newHistory string) (*dynamodb.UpdateItemOutput, error) {
	dynamoLibs := aw.GetDynamoDB()
	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(ID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":valhistory": {
				SS: aws.StringSlice([]string{
					newHistory,
					oldHistory,
				}),
			},
			":valstatusText": {
				S: aws.String(status),
			},
		},

		TableName:        aws.String(os.Getenv("AWS_DYNAMO_TABLE")),
		ReturnValues:     aws.String("ALL_NEW"),
		UpdateExpression: aws.String("SET history = :valhistory, statusText = :valstatusText"),
	}
	result, err := dynamoLibs.UpdateItem(input)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetDynamoData ..
func (aw *Aws) GetDynamoData(ID string) (*dynamodb.GetItemOutput, error) {
	dynamoLibs := aw.GetDynamoDB()
	result, err := dynamoLibs.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("AWS_DYNAMO_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(ID),
			},
		},
		// ProjectionExpression: expr.Projection(),
	})
	return result, err
}

// GetDynamoHistory ..
func (aw *Aws) GetDynamoHistory(receiverAddress string) (*dynamodb.ScanOutput, error) {
	dynamoLibs := aw.GetDynamoDB()
	filter := expression.Name("receiverAddress").Equal(expression.Value(receiverAddress))
	proj := expression.NamesList(
		expression.Name("id"),
		expression.Name("history"),
		expression.Name("data"),
		expression.Name("receiverAddress"),
		expression.Name("createdAt"),
		expression.Name("statusText"),
		expression.Name("type"),
		expression.Name("updatedAt"),
	)
	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(proj).Build()
	if err != nil {
		return nil, err
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(os.Getenv("AWS_DYNAMO_TABLE")),
	}

	result, err := dynamoLibs.Scan(params)
	if err != nil {
		return nil, err
	}
	return result, nil
}
