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
