package gateway

import (
	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
)

// Gateway ...
type Gateway struct{}

// GatewayHandler Handler
func GatewayHandler() *Gateway {
	return &Gateway{}
}

//GatewayInterface declare All Method
type GatewayInterface interface {
	PostNotification() *entity.PostNotificationResponse
}

// PostNotification ...
// return *entity.PostNotificationResponse
func (gateway *Gateway) PostNotification() *entity.PostNotificationResponse {

}
