package provider

import (
	"encoding/json"
	"log"
	"os"

	entity "github.com/sofyan48/otp/src/entity/http/v1"
)

// InterceptorMessages ...
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
	operator, msisdn := prv.OperatorChecker(data.Payload.Msisdn)
	historyPayload := &entity.PayloadPostNotificationRequest{}
	historyPayload.Msisdn = msisdn
	historyPayload.OTP = data.Payload.OTP
	historyPayload.Text = data.Payload.Text
	historyValue := &entity.HistoryItem{}
	historyValue.CallbackData = itemDynamo.ID
	historyValue.Payload = historyPayload
	historyValue.Response = "interceptors"
	if operator.Name == "xl" {
		historyValue.Provider = dataThirdParty[1].Provider
		history := map[string]*entity.HistoryItem{
			dataThirdParty[1].Provider: historyValue,
		}
		itemDynamo.History = history

	} else {
		historyValue.Provider = dataThirdParty[0].Provider
		history := map[string]*entity.HistoryItem{
			dataThirdParty[0].Provider: historyValue,
		}
		itemDynamo.History = history
	}
	return itemDynamo

}

// InterceptorEmail ...
func (prv *Providers) InterceptorEmail(data *entity.PostNotificationRequestEmail) *entity.DynamoItemEmail {
	itemDynamo := &entity.DynamoItemEmail{}
	marshalData, err := json.Marshal(data.Payload.Data)
	itemDynamo.Data = string(marshalData)
	itemDynamo.ReceiverAddress = data.Payload.To

	itemDynamo.StatusText = "QUEUE"
	itemDynamo.ID = data.UUID
	itemDynamo.Type = data.Type

	dataThirdParty := make([]entity.DataProvider, 0)
	err = json.Unmarshal([]byte(os.Getenv("EMAIL_ORDER_CONF")), &dataThirdParty)
	if err != nil {
		log.Println(err)
	}

	historyPayload := &entity.PayloadRequestEmail{}
	historyPayload.Data = data.Payload.Data
	historyPayload.To = data.Payload.To
	historyPayload.From = data.Payload.From

	historyValue := &entity.EmailHistoryItem{}
	historyValue.Payload = historyPayload
	historyValue.Response = "interceptors"
	historyValue.Provider = dataThirdParty[0].Provider
	history := map[string]*entity.EmailHistoryItem{
		dataThirdParty[0].Provider: historyValue,
	}
	itemDynamo.History = history
	return itemDynamo

}

// InfobipSender ..
func (prv *Providers) InfobipSender() {

}

// WavecellSender ...
func (prv *Providers) WavecellSender() {

}
