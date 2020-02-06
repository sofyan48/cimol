package users

import (
	"github.com/gin-gonic/gin"

	ctrlNotif "github.com/sofyan48/rll-daemon-new/src/controller/v1/gateway"
	svcNotif "github.com/sofyan48/rll-daemon-new/src/service/v1/gateway"
	"github.com/sofyan48/rll-daemon-new/src/util/middleware"
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
		ServiceGateway: *svcNotif.GatewayHandler(),
	}

	// //********* Calling Handler To Routers *********//
	rLoader.routerPostNotification(router, postNotifHandler)

}

//********* Routing API *********//

// routerDefinition Routes
// @router: gin Engine
// @handler: ControllerGateway
func (rLoader *V1RouterLoader) routerPostNotification(router *gin.Engine, handler *ctrlNotif.ControllerGateway) {
	group := router.Group("/v1/notification")
	group.POST("", handler.PostNotification)
	group.GET("history/:receiverAddress", handler.GetHistory)
	group.GET("id/:id", handler.GetByID)
}
