package receivers

import (
	"github.com/sofyan48/cimol/src/controller/v1/receivers/notification"
)

// NotificationReceiver ...
type ControllerReceiver struct {
	Receivers notification.NotificationReceiverInterface
}
