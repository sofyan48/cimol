package receivers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
	service "github.com/sofyan48/rll-daemon-new/src/service/v1/receivers"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/rest"
)

// ControllerReceiver ...
type ControllerReceiver struct {
	ServiceReceivers service.ReceiverInterface
}

// InfobipReceiver ...
func (ctrl *ControllerReceiver) InfobipReceiver(context *gin.Context) {
	payload := &entity.InfobipCallBackRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, payload)
}

// WavecellReceiver ...
func (ctrl *ControllerReceiver) WavecellReceiver(context *gin.Context) {
	payload := &entity.WavecelllCallBackRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, payload)
}
