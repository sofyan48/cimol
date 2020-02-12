package transmiter

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
)

func (trs *Transmiter) intercepActionShard(data *entity.DynamoItem) {
	dataThirdParty := make([]entity.DataProvider, 0)
	err := json.Unmarshal([]byte(os.Getenv("SMS_ORDER_CONF")), &dataThirdParty)
	if err != nil {
		log.Println(err)
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
			break
		}
	}
	fmt.Println("Provider: ", historyProvider)
	fmt.Println("History: ", history)
	// trs.wavecellActionShard(historyProvider, history)

}
