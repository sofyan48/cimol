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
	payload := &entity.InfobipCallbackRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, err.Error())
		return
	}
	ctrl.ServiceReceivers.InfobipReceiver(payload.Results[0].CallbackData, payload)
	rest.ResponseData(context, http.StatusOK, payload)
	return
}

// WavecellReceiver ...
func (ctrl *ControllerReceiver) WavecellReceiver(context *gin.Context) {
	payload := &entity.WavecelllCallBackRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, err.Error())
		return
	}
	id := "afc3651c-ff8e-4d07-83cc-433a2e67d775"
	ctrl.ServiceReceivers.WavecellReceiver(id, payload)
	rest.ResponseData(context, http.StatusOK, payload)
	return
}
