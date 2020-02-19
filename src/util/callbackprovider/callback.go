package callbackprovider

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
	"github.com/sofyan48/cimol/src/util/helper/libaws"
	"github.com/sofyan48/cimol/src/util/helper/provider"
)

// ProviderCallback ...
type ProviderCallback struct {
	AwsLib   libaws.AwsInterface
	Provider provider.ProvidersInterface
}

// ProviderCallbackHandler ...
func ProviderCallbackHandler() *ProviderCallback {
	return &ProviderCallback{
		AwsLib:   libaws.AwsHAndler(),
		Provider: provider.ProvidersHandler(),
	}
}

// ProviderCallbackInterface ...
type ProviderCallbackInterface interface {
	InfobipCallback(dynamo *entity.DynamoItemResponse, data *entity.InfobipCallbackRequest, history *entity.HistoryItem)
	WavecellCallback(dynamo *entity.DynamoItemResponse, data *entity.WavecellCallBackRequest, history *entity.HistoryItem)
	wavecellMessagesNotSuccess(dynamo *entity.DynamoItemResponse, data *entity.WavecellCallBackRequest, history *entity.HistoryItem)
	infobipMessagesNotSuccess(dynamo *entity.DynamoItemResponse, data *entity.WavecellCallBackRequest, history *entity.HistoryItem)
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

func (callback *ProviderCallback) infobipMessagesNotSuccess(dynamo *entity.DynamoItemResponse,
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
