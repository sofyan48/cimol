package email

import (
	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

// SengridSend ..
func (email *EmailModuls) SengridSend(data *entity.EmailHistoryItem) {
	email.Sendgrid.SendEmail(data)
}
