package notification

import (
	"sync"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
	"github.com/sofyan48/cimol/src/util/helper/libaws"
	"github.com/sofyan48/cimol/src/util/helper/libsendgrid"
	"github.com/sofyan48/cimol/src/util/helper/provider"
)

// NotificationGateway ...
type NotificationGateway struct {
	AwsLib    libaws.AwsInterface
	Providers provider.ProvidersInterface
	Sendgrid  libsendgrid.LibsendgridInterface
}

// NotificationGatewayHandler Handler
func NotificationGatewayHandler() *NotificationGateway {
	return &NotificationGateway{
		AwsLib:    libaws.AwsHAndler(),
		Providers: provider.ProvidersHandler(),
		Sendgrid:  libsendgrid.LibSendgridHandler(),
	}
}

// NotificationGatewayInterface ...
type NotificationGatewayInterface interface {
	GetByID(ID string) (*entity.DynamoItemHistory, error)
	GetHistory(msisdn string) ([]entity.DynamoItemHistory, error)
	PostNotificationEmail(data *entity.PostNotificationRequestEmail, wg *sync.WaitGroup)
	PostNotification(data *entity.PostNotificationRequest, wg *sync.WaitGroup)
	PostNotificationPush()
}
