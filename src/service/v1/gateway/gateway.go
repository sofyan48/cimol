package gateway

import (
	"github.com/sofyan48/cimol/src/service/v1/gateway/notification"
)

// Gateway ...
type Gateway struct {
	Notification notification.NotificationGatewayInterface
}

// GatewayHandler Handler
func GatewayHandler() *Gateway {
	return &Gateway{}
}

//GatewayInterface declare All Method
type GatewayInterface interface {
	GetNotification() *notification.NotificationGateway
}

func (gw *Gateway) GetNotification() *notification.NotificationGateway {
	return notification.NotificationGatewayHandler()
}
