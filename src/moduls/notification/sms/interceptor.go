package sms

import (
	"encoding/json"
	"log"
	"os"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
	"github.com/sofyan48/cimol/src/util/helper/libaws"
	"github.com/sofyan48/cimol/src/util/logging"
	"github.com/sofyan48/cimol/src/util/provider"
	"github.com/sofyan48/cimol/src/util/request"
)

type SMSModuls struct {
	AwsLibs   libaws.AwsInterface
	Logs      logging.LogInterface
	Provider  provider.ProvidersInterface
	Requester request.RequesterInterface
}

func SMSModulsHandler() *SMSModuls {
	return &SMSModuls{
		AwsLibs:   libaws.AwsHAndler(),
		Logs:      logging.LogHandler(),
		Provider:  provider.ProvidersHandler(),
		Requester: request.RequesterHandler(),
	}
}

type SMSModulsInterface interface {
	IntercepActionShardSMS(data *entity.DynamoItem)
	InfobipActionShardOTP(history string, payload *entity.HistoryItem)
	WavecellActionShard(history string, payload *entity.HistoryItem)
}

// IntercepActionShardSMS ...
func (sms *SMSModuls) IntercepActionShardSMS(data *entity.DynamoItem) {
	dataThirdParty := make([]entity.DataProvider, 0)
	err := json.Unmarshal([]byte(os.Getenv("SMS_ORDER_CONF")), &dataThirdParty)
	if err != nil {
		sms.Logs.Write("Transmitter SMS", err.Error())
	}
	history := &entity.HistoryItem{}
	var historyProvider string
	for _, i := range dataThirdParty {
		if data.History[i.Provider] != nil {
			history.CallbackData = data.History[i.Provider].CallbackData
			history.DeliveryReport = data.History[i.Provider].DeliveryReport
			history.Payload = data.History[i.Provider].Payload
			history.Response = data.History[i.Provider].Response
			historyProvider = data.History[i.Provider].Provider
			history.Provider = data.History[i.Provider].Provider
			break
		}
	}
	switch historyProvider {
	case "infobip":
		sms.InfobipActionShardOTP(historyProvider, history)
	case "wavecell":
		sms.WavecellActionShard(historyProvider, history)
	case "gosms":
		log.Println("INCOMING")
	case "twilio":
		log.Println("INCOMING")
	default:
		sms.InfobipActionShardOTP(historyProvider, history)
	}

}

func checkEnvironment() bool {
	envi := os.Getenv("APP_ENVIRONMENT")
	return envi == "development" || envi == "staging"
}

// updateDynamoTransmitt ...
func (sms *SMSModuls) updateDynamoTransmitt(ID, status, data string, history *entity.HistoryItem) (string, error) {
	result, err := sms.AwsLibs.UpdateDynamo(ID, status, data, history)
	return result.GoString(), err
}
