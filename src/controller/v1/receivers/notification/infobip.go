package notification

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	entity "github.com/sofyan48/cimol/src/entity/http/v1"
	"github.com/sofyan48/cimol/src/util/helper/rest"
)

// InfobipReceiver ...
func (notif *NotificationReceiver) InfobipReceiver(context *gin.Context) {
	payload := &entity.InfobipCallbackRequest{}
	err := context.ShouldBind(payload)
	if err != nil {
		rest.ResponseMessages(context, http.StatusBadRequest, err.Error())
		return
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go notif.ServiceReceivers.InfobipReceiver(payload.Results[0].CallbackData, payload)

	rest.ResponseMessages(context, http.StatusOK, "Success")

	return
}
