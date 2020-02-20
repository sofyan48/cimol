package receivers

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

// WavecellReceiver ...
func (rcv *Receiver) WavecellReceiver(ID string, data *entity.WavecellCallBackRequest) (string, error) {
	dynamoItem := &entity.DynamoItemResponse{}
	dynamoData, err := rcv.AwsLib.GetDynamoData(ID)
	if err != nil {
		rcv.Logs.Write("Receiver", err.Error())
	}
	err = dynamodbattribute.UnmarshalMap(dynamoData.Item, &dynamoItem)
	if err != nil {
		rcv.Logs.Write("Receiver", err.Error())
	}

	historyItems := &entity.HistoryItem{}
	err = json.Unmarshal([]byte(dynamoItem.History[1]), historyItems)
	if err != nil {
		rcv.Logs.Write("Receiver", err.Error())
	}
	rcv.Callback.WavecellCallback(dynamoItem, data, historyItems)
	return "", nil
}
