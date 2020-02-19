package users

import (
	"github.com/gin-gonic/gin"

	ctrlNotif "github.com/sofyan48/cimol/src/controller/v1/gateway"
	svcNotif "github.com/sofyan48/cimol/src/service/v1/gateway"

	ctrlRcv "github.com/sofyan48/cimol/src/controller/v1/receivers"
	svcRcv "github.com/sofyan48/cimol/src/service/v1/receivers"
	"github.com/sofyan48/cimol/src/util/middleware"
)

// V1RouterLoader types
type V1RouterLoader struct {
	Middleware middleware.DefaultMiddleware
}

// V1Router Params
// @router: gin.Engine
func (rLoader *V1RouterLoader) V1Router(router *gin.Engine) {

	// post Notif Handler Routes
	postNotifHandler := &ctrlNotif.ControllerGateway{
		ServiceGateway: svcNotif.GatewayHandler(),
	}

	// receivers
	rcvHandler := &ctrlRcv.ControllerReceiver{
		ServiceReceivers: svcRcv.ReceiverHandler(),
	}

	// //********* Calling Handler To Routers *********//
	rLoader.routerPostNotification(router, postNotifHandler)
	rLoader.routerReceiver(router, rcvHandler)

}

//********* Routing API *********//

// routerDefinition Routes
// @router: gin Engine
// @handler: ControllerGateway
func (rLoader *V1RouterLoader) routerPostNotification(router *gin.Engine, handler *ctrlNotif.ControllerGateway) {
	group := router.Group("/v1/notification")
	group.POST("sms", handler.PostNotification)
	group.POST("email", handler.PostNotificationEmail)
	group.POST("push", handler.PostNotificationPush)
	group.GET("history/:receiverAddress", handler.GetHistory)
	group.GET("id/:id", handler.GetByID)
}

// routerDefinition Routes
// @router: gin Engine
// @handler: ControllerReceiver
func (rLoader *V1RouterLoader) routerReceiver(router *gin.Engine, handler *ctrlRcv.ControllerReceiver) {
	group := router.Group("/v1/receiver")
	group.POST("infobip", handler.InfobipReceiver)
	group.POST("wavecell", handler.WavecellReceiver)
}
