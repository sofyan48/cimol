package transmiter

import (
	entity "github.com/sofyan48/otp/src/entity/http/v1"
)

// SengridSend ..
func (trs *Transmiter) SengridSend(data *entity.EmailHistoryItem) {
	trs.Sendgrid.SendEmail(data)
}
