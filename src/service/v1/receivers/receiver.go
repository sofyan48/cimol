package receivers

import (
	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/libaws"
)

// Receiver ...
type Receiver struct {
	AwsLib libaws.AwsInterface
}

// ReceiverHandler ...
func ReceiverHandler() *Receiver {
	return &Receiver{
		AwsLib: libaws.AwsHAndler(),
	}
}

// ReceiverInterface ...
type ReceiverInterface interface {
	InfobipReceiver(data *entity.InfobipCallBackRequest)
	WavecellReceiver(data *entity.WavecelllCallBackRequest)
}

// InfobipReceiver ...
func (rcv *Receiver) InfobipReceiver(data *entity.InfobipCallBackRequest) {

}

// WavecellReceiver ...
func (rcv *Receiver) WavecellReceiver(data *entity.WavecelllCallBackRequest) {

}
