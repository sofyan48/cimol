package callbackprovider

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"sync"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

// WavecellCallback ...
func (callback *ProviderCallback) WavecellCallback(dynamo *entity.DynamoItemResponse,
	data *entity.WavecellCallBackRequest, history *entity.HistoryItem) {
	if validateStatusWavecell(data.Status) {
		callback.wavecellMessagesNotSuccess(dynamo, data, history)
		return
	}
	callback.wavecellSuccessReport(dynamo, data, history)
}

func (callback *ProviderCallback) wavecellSuccessReport(dynamo *entity.DynamoItemResponse,
	data *entity.WavecellCallBackRequest, history *entity.HistoryItem) {
	oldHistory, _ := json.Marshal(dynamo.History)
	newHistory, _ := json.Marshal(data)
	callback.AwsLib.CallbackSendUpdate(history.CallbackData, data.Status, string(oldHistory), string(newHistory))
}

func (callback *ProviderCallback) wavecellMessagesNotSuccess(dynamo *entity.DynamoItemResponse,
	data *entity.WavecellCallBackRequest, history *entity.HistoryItem) {
	dataThirdParty := make([]entity.DataProvider, 0)
	err := json.Unmarshal([]byte(os.Getenv("SMS_ORDER_CONF")), &dataThirdParty)
	if err != nil {
		log.Println(err)
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
		dataThirdParty[0].Provider: historyValue,
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

func validateStatusWavecell(status string) bool {
	statusData := []string{
		"REJECTED BY CARRIER",
		"REJECTED BY DEVICE",
		"REJECTED BY WAVECELL",
	}
	for _, i := range statusData {
		if strings.EqualFold(i, status) {
			return true
		}
	}
	return false
}
