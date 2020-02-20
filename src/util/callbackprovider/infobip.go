package callbackprovider

import (
	"encoding/json"
	"os"
	"strings"
	"sync"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

// InfobipCallback ...
func (callback *ProviderCallback) InfobipCallback(dynamo *entity.DynamoItemResponse,
	data *entity.InfobipCallbackRequest, history *entity.HistoryItem) {
	if validateStatusInfobip(data.Results[0].Status.Name) {
		callback.InfobipCallback(dynamo, data, history)
		return
	}
	wg := &sync.WaitGroup{}
	go callback.infobipSuccessReport(dynamo, data, history, wg)
}

func (callback *ProviderCallback) infobipSuccessReport(dynamo *entity.DynamoItemResponse,
	data *entity.InfobipCallbackRequest, history *entity.HistoryItem, wg *sync.WaitGroup) {
	oldHistory, _ := json.Marshal(dynamo.History)
	newHistory, _ := json.Marshal(data)
	callback.AwsLib.CallbackSendUpdate(history.CallbackData, data.Results[0].Status.Name, string(oldHistory), string(newHistory))
	wg.Done()
}

func (callback *ProviderCallback) infobipMessagesNotSuccess(dynamo *entity.DynamoItemResponse,
	data *entity.WavecellCallBackRequest, history *entity.HistoryItem) {
	dataThirdParty := make([]entity.DataProvider, 0)
	err := json.Unmarshal([]byte(os.Getenv("SMS_ORDER_CONF")), &dataThirdParty)
	if err != nil {
		callback.Logs.Write("Callback", err.Error())
	}
	historyPayload := &entity.PayloadPostNotificationRequest{}
	_, msisdn := callback.Provider.OperatorChecker(dynamo.ReceiverAddress)
	historyPayload.Msisdn = msisdn
	historyPayload.OTP = history.Payload.OTP
	historyPayload.Text = history.Payload.Text

	historyValue := &entity.HistoryItem{}
	historyValue.CallbackData = dynamo.ID
	historyValue.Payload = historyPayload
	historyValue.Response = "interceptors"
	historyValue.Provider = dataThirdParty[1].Provider

	itemDynamo := &entity.DynamoItem{}
	itemDynamo.ID = dynamo.ID
	itemDynamo.Data = dynamo.Data
	itemDynamo.History = map[string]*entity.HistoryItem{
		dataThirdParty[1].Provider: historyValue,
	}

	itemDynamo.ReceiverAddress = dynamo.ReceiverAddress
	itemDynamo.StatusText = "RELOADING"
	itemDynamo.Type = dynamo.Type
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go callback.AwsLib.InputDynamo(itemDynamo, wg)
	wg.Add(1)
	go callback.AwsLib.SendStart(itemDynamo.ID, itemDynamo, "interceptors", wg)
}

func validateStatusInfobip(status string) bool {
	statusData := []string{
		"UNDELIVERABLE_REJECTED_OPERATOR",
		"UNDELIVERABLE_NOT_DELIVERED",
		"PENDING_ENROUTE",
	}
	for _, i := range statusData {
		if strings.EqualFold(i, status) {
			return true
		}
	}
	return false
}
