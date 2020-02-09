package provider

import (
	"encoding/json"
	"log"
	"os"

	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
)

func (prv *Providers) InterceptorMessages(data *entity.PostNotificationRequest) *entity.DynamoItem {
	itemDynamo := &entity.DynamoItem{}
	itemDynamo.Data = data.Payload.Text
	itemDynamo.ReceiverAddress = data.Payload.Msisdn

	itemDynamo.StatusText = "QUEUE"
	itemDynamo.ID = data.UUID
	itemDynamo.Type = data.Type

	dataThirdParty := make([]entity.DataProvider, 0)
	err := json.Unmarshal([]byte(os.Getenv("SMS_ORDER_CONF")), &dataThirdParty)
	if err != nil {
		log.Println(err)
	}

	operator := prv.OperatorChecker(data.Payload.Msisdn)
	if operator.Name == "xl" {
		history := map[string]*entity.HistoryItem{
			"wavecell": &entity.HistoryItem{},
		}
		itemDynamo.History = history
	}

	return itemDynamo

}

// InfobipSender ..
func (prv *Providers) InfobipSender() {

}

// WavecellSender ...
func (prv *Providers) WavecellSender() {

}
