package callbackprovider

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	entity "github.com/sofyan48/otp/src/entity/http/v1"
)

// InfobipCallback ...
func (callback *ProviderCallback) InfobipCallback(dynamo *entity.DynamoItemResponse,
	data *entity.InfobipCallbackRequest, history *entity.HistoryItem) {
	if validateStatusInfobip(data.Results[0].Status.Name) {
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

	} else {
		fmt.Println("UPDATE DATA")
	}
}

func validateStatusInfobip(status string) bool {
	statusData := []string{
		"UNDELIVERABLE_REJECTED_OPERATOR",
		"UNDELIVERABLE_NOT_DELIVERED",
		"PENDING_ENROUTE",
		"PENDING_ACCEPTED",
	}
	for _, i := range statusData {
		if strings.EqualFold(i, status) {
			return true
		}
	}
	return false
}
