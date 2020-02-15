package libsendgrid

import (
	"encoding/json"
	"os"

	entity "github.com/sofyan48/otp/src/entity/http/v1"
	"github.com/sofyan48/otp/src/util/helper/request"
)

// Libsendgrid ...
type Libsendgrid struct {
	Requester request.RequesterInterface
}

// LibSendgridHandler ...
func LibSendgridHandler() *Libsendgrid {
	return &Libsendgrid{
		Requester: request.RequesterHandler(),
	}
}

// LibsendgridInterface ...
type LibsendgridInterface interface {
	GetTemplateID(ID string) (*entity.TemplateResponse, error)
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

func (libsengrid *Libsendgrid) SendEmail() {

}
