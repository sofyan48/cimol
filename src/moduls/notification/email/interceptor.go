package email

import (
	"encoding/json"
	"fmt"
	"os"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
	"github.com/sofyan48/cimol/src/util/helper/libsendgrid"
	"github.com/sofyan48/cimol/src/util/helper/logging"
)

// EmailModuls ...
type EmailModuls struct {
	Logs     logging.LogInterface
	Sendgrid libsendgrid.LibsendgridInterface
}

// EmailModulsHandler ...
func EmailModulsHandler() *EmailModuls {
	return &EmailModuls{
		Logs:     logging.LogHandler(),
		Sendgrid: libsendgrid.LibSendgridHandler(),
	}
}

// EmailTransmiterInterface ...
type EmailTransmiterInterface interface {
	IntercepActionShardEmail(data *entity.DynamoItemEmail)
	SengridSend(data *entity.EmailHistoryItem)
}

// IntercepActionShardEmail ...
func (email *EmailModuls) IntercepActionShardEmail(data *entity.DynamoItemEmail) {
	dataThirdParty := make([]entity.DataProvider, 0)
	err := json.Unmarshal([]byte(os.Getenv("EMAIL_ORDER_CONF")), &dataThirdParty)
	if err != nil {
		email.Logs.Write("Transmitter Email", err.Error())
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
		email.SengridSend(history)
	default:
		fmt.Println("Coming Soon")
	}
}
