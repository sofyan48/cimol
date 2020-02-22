package notification

import (
	"sync"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

// PostNotificationEmail ...
func (gateway *NotificationGateway) PostNotificationEmail(data *entity.PostNotificationRequestEmail, wg *sync.WaitGroup) {
	itemDynamo := gateway.Providers.InterceptorEmail(data)
	wg.Add(1)
	gateway.AwsLib.SendMail(data.UUID, itemDynamo, "email", wg)
	wg.Add(1)
	go gateway.AwsLib.InputDynamoEmail(itemDynamo, wg)

}
