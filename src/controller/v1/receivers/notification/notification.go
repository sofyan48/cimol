package notification

import (
	"github.com/gin-gonic/gin"
	service "github.com/sofyan48/cimol/src/service/v1/receivers"
)

// NotificationReceiver ...
type NotificationReceiver struct {
	ServiceReceivers service.ReceiverInterface
}

// NotificationReceiverHandler ...
func NotificationReceiverHandler() *NotificationReceiver {
	return &NotificationReceiver{
		ServiceReceivers: service.ReceiverHandler(),
	}
}

type NotificationReceiverInterface interface {
	InfobipReceiver(context *gin.Context)
	WavecellReceiver(context *gin.Context)
}
