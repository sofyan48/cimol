package callbackprovider

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
)

// InfobipCallback ...
func (callback *ProviderCallback) InfobipCallback(dynamo *entity.DynamoItemResponse,
	data *entity.InfobipCallbackRequest) {
	fmt.Println(validateStatusInfobip(data.Results[0].Status.Name))
	if validateStatusInfobip(data.Results[0].Status.Name) {
		statusHistory, err := json.Marshal(data.Results[0])
		if err != nil {
			log.Println("Error", err)
		}
		dataThirdParty := make([]entity.DataProvider, 0)
		err = json.Unmarshal([]byte(os.Getenv("SMS_ORDER_CONF")), &dataThirdParty)
		if err != nil {
			log.Println(err)
		}
		itemDynamo := &entity.DynamoItem{}
		historyPayload := &entity.PayloadPostNotificationRequest{}
		historyPayload.Msisdn = fmt.Sprintf("%v", dynamo.ReceiverAddress)
		historyPayload.OTP = true
		historyPayload.Text = fmt.Sprintf("%v", dynamo.Data)

		historyValue := &entity.HistoryItem{}
		historyValue.CallbackData = itemDynamo.ID
		historyValue.Payload = historyPayload
		historyValue.Response = "interceptors"

		itemDynamo.Data = fmt.Sprintf("%v", dynamo.Data)
		itemDynamo.History = map[string]*entity.HistoryItem{
			dataThirdParty[1].Provider: historyValue,
		}
		itemDynamo.ID = fmt.Sprintf("%v", dynamo.ID)
		itemDynamo.ReceiverAddress = fmt.Sprintf("%v", dynamo.ReceiverAddress)
		itemDynamo.StatusText = "QUEUE"
		itemDynamo.Type = "sms"

		_, err = callback.AwsLib.UpdateDynamo(data.Results[0].CallbackData,
			data.Results[0].Status.Name, string(statusHistory), historyValue)
		if err != nil {
			log.Println("Error", err)
		}
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go callback.AwsLib.SendStart(itemDynamo.ID, itemDynamo, "interceptors", wg)
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
			break
		}
	}
	return false
}
