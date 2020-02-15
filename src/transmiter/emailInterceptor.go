package transmiter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	entity "github.com/sofyan48/otp/src/entity/http/v1"
)

func (trs *Transmiter) intercepActionShardEmail(data *entity.DynamoItemEmail) {
	dataThirdParty := make([]entity.DataProvider, 0)
	err := json.Unmarshal([]byte(os.Getenv("EMAIL_ORDER_CONF")), &dataThirdParty)
	if err != nil {
		log.Println(err)
	}
	history := &entity.EmailHistoryItem{}
	var historyProvider string
	for _, i := range dataThirdParty {
		if data.History[i.Provider] != nil {
			history.DeliveryReport = data.History[i.Provider].DeliveryReport
			history.Payload = data.History[i.Provider].Payload
			history.Response = data.History[i.Provider].Response
			historyProvider = data.History[i.Provider].Provider
			history.Provider = data.History[i.Provider].Provider
			break
		}
	}
	switch historyProvider {
	case "sendgrid":
		fmt.Println("Coming Soon", history)
	default:
		fmt.Println("Coming Soon")
	}

}
