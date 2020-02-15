package libsendgrid

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	entity "github.com/sofyan48/otp/src/entity/http/v1"
	"github.com/sofyan48/otp/src/util/helper/mailtrap"
	"github.com/sofyan48/otp/src/util/helper/request"
)

// Libsendgrid ...
type Libsendgrid struct {
	Requester request.RequesterInterface
	Mailtrap  mailtrap.MailtrapInterface
}

// LibSendgridHandler ...
func LibSendgridHandler() *Libsendgrid {
	return &Libsendgrid{
		Requester: request.RequesterHandler(),
		Mailtrap:  mailtrap.MailtrapHandler(),
	}
}

// LibsendgridInterface ...
type LibsendgridInterface interface {
	GetTemplateID(ID string) (*entity.TemplateResponse, error)
	SendEmail(data *entity.EmailHistoryItem)
}

// GetTemplateID ...
func (libsengrid *Libsendgrid) GetTemplateID(ID string) (*entity.TemplateResponse, error) {
	results := &entity.TemplateResponse{}
	uri := os.Getenv("SENDGRID_URL") + "/v3/templates/" + ID
	auth := os.Getenv("SENDGRID_AUTH_TYPE") + " " + os.Getenv("SENDGRID_TOKEN")
	data, err := libsengrid.Requester.GET(uri, auth)
	if err != nil {
		return results, err
	}
	err = json.Unmarshal(data, results)
	return results, err

}

// SendEmail ...
func (libsengrid *Libsendgrid) SendEmail(history *entity.EmailHistoryItem) {
	if os.Getenv("APP_ENVIRONMENT") != "production" {
		templateData, err := libsengrid.GetTemplateID(history.Payload.TemplateID)
		if err != nil {
			log.Println("Error: ", err)
		}
		htmlContent := templateData.Versions[0].HTMLContent
		for key, word := range history.Payload.Data {
			htmlContent = strings.Replace(htmlContent, key, word, -1)
		}
		err = libsengrid.Mailtrap.SendMail(history.Payload.To, history.Payload.Subject, htmlContent)
		if err != nil {
			log.Println("Error Sending Email: ", err)
		}
		return
	}

}
