package gateway

import (
	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/libaws"
)

// Gateway ...
type Gateway struct {
	AwsLib *libaws.Aws
}

// GatewayHandler Handler
func GatewayHandler() *Gateway {
	return &Gateway{
		AwsLib: libaws.AwsHAndler(),
	}
}

//GatewayInterface declare All Method
type GatewayInterface interface {
	PostNotification() *entity.PostNotificationResponse
}

// PostNotification ...
// return *entity.PostNotificationResponse
func (gateway *Gateway) PostNotification() *entity.PostNotificationResponse {
	result := &entity.PostNotificationResponse{}
	result.ID = "ID"
	result.Status = "Status"
	return result
}
