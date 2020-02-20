package transmiter

import (
	"encoding/json"
	"log"
	"os"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

func (trs *Transmiter) intercepActionShardSMS(data *entity.DynamoItem) {
	dataThirdParty := make([]entity.DataProvider, 0)
	err := json.Unmarshal([]byte(os.Getenv("SMS_ORDER_CONF")), &dataThirdParty)
	if err != nil {
		trs.Logs.Write("Transmitter SMS", err.Error())
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
		trs.infobipActionShardOTP(historyProvider, history)
	case "wavecell":
		trs.wavecellActionShard(historyProvider, history)
	case "gosms":
		log.Println("INCOMING")
	case "twilio":
		log.Println("INCOMING")
	default:
		trs.infobipActionShardOTP(historyProvider, history)
	}

}
