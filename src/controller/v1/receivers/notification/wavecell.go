package notification

import (
	"net/http"

	"github.com/gin-gonic/gin"
	entity "github.com/sofyan48/cimol/src/entity/http/v1"
	"github.com/sofyan48/cimol/src/util/helper/rest"
)

// WavecellReceiver ...
func (notif *NotificationReceiver) WavecellReceiver(context *gin.Context) {
	payload := &entity.WavecellCallBackRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, err.Error())
		return
	}
	id := payload.ClientMessageID
	notif.ServiceReceivers.WavecellReceiver(id, payload)
	rest.ResponseMessages(context, http.StatusOK, "QUEUE")
	return
}
