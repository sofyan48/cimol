package receivers

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	entity "github.com/sofyan48/cimol/src/entity/http/v1"
	service "github.com/sofyan48/cimol/src/service/v1/receivers"
	"github.com/sofyan48/cimol/src/util/helper/rest"
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
	wg := sync.WaitGroup{}
	wg.Add(1)
	go ctrl.ServiceReceivers.InfobipReceiver(payload.Results[0].CallbackData, payload)

	rest.ResponseMessages(context, http.StatusOK, "Success")

	return
}

// WavecellReceiver ...
func (ctrl *ControllerReceiver) WavecellReceiver(context *gin.Context) {
	payload := &entity.WavecellCallBackRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, err.Error())
		return
	}
	id := payload.ClientMessageID
	ctrl.ServiceReceivers.WavecellReceiver(id, payload)
	rest.ResponseMessages(context, http.StatusOK, "QUEUE")
	return
}
