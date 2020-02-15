package mailtrap

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"os"

	"github.com/sofyan48/otp/src/util/helper/request"
)

// Mailtrap ...
type Mailtrap struct {
	Requester request.RequesterInterface
}

// MailtrapHandler ...
func MailtrapHandler() *Mailtrap {
	return &Mailtrap{
		Requester: request.RequesterHandler(),
	}
}

// MailtrapInterface ...
type MailtrapInterface interface {
	SendMail(data string)
}

// SendMail ...
func (trap *Mailtrap) SendMail(data string) {
	auth := smtp.PlainAuth("", os.Getenv("MAILTRAP_USERNAME"), os.Getenv("MAILTRAP_PASSWORD"), os.Getenv("MAILTRAP_HOST"))
	to := []string{"meongbego@gmail.com"}
	from := os.Getenv("MAILTRAP_IDENTITY")
	addr := os.Getenv("MAILTRAP_HOST") + ":" + os.Getenv("MAILTRAP_PORT")
	msg, _ := json.Marshal(data)
	err := smtp.SendMail(addr, auth, from, to, msg)
	if err != nil {
		fmt.Println(err.Error())
	}
}
