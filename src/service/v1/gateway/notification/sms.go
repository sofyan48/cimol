package notification

import (
	"sync"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

// PostNotification ...
func (gateway *NotificationGateway) PostNotification(data *entity.PostNotificationRequest, wg *sync.WaitGroup) {
	itemDynamo := gateway.Providers.InterceptorMessages(data)
	wg.Add(1)
	go gateway.AwsLib.SendStart(data.UUID, itemDynamo, "interceptors", wg)
	wg.Add(1)
	go gateway.AwsLib.InputDynamo(itemDynamo, wg)

}
