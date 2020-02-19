package receivers

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

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
