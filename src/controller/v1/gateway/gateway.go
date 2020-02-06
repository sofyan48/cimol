package gateway

import (
	"net/http"

	"github.com/gin-gonic/gin"
	service "github.com/sofyan48/rll-daemon-new/src/service/v1/gateway"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/rest"
)

// ControllerGateway ...
type ControllerGateway struct {
	ServiceGateway service.Gateway
}

// PostNotification ...
func (ctrl *ControllerGateway) PostNotification(context *gin.Context) {
	postNotification := ctrl.ServiceGateway.PostNotification()
	rest.ResponseData(context, http.StatusOK, postNotification)
}

// GetHistory ...
func (ctrl *ControllerGateway) GetHistory(context *gin.Context) {
	history, err := ctrl.ServiceGateway.GetHistory(context.Param("receiverAddress"))
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, history)
	return
}

// GetByID ...
func (ctrl *ControllerGateway) GetByID(context *gin.Context) {
	data, err := ctrl.ServiceGateway.GetByID(context.Param("id"))
	if err != nil {
		rest.ResponseMessages(context, http.StatusInternalServerError, err.Error())
		return
	}
	rest.ResponseData(context, http.StatusOK, data)
	return
}
