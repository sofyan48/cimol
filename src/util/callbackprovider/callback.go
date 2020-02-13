package callbackprovider

import (
	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
	"github.com/sofyan48/rll-daemon-new/src/util/helper/libaws"
)

// ProviderCallback ...
type ProviderCallback struct {
	AwsLib libaws.AwsInterface
}

// ProviderCallbackHandler ...
func ProviderCallbackHandler() *ProviderCallback {
	return &ProviderCallback{
		AwsLib: libaws.AwsHAndler(),
	}
}

// ProviderCallbackInterface ...
type ProviderCallbackInterface interface {
	SendToBroker(data *entity.DynamoItem)
	InfobipCallback(dynamo *entity.DynamoItemResponse, data *entity.InfobipCallbackRequest)
}

// SendToBroker ...
func (callback *ProviderCallback) SendToBroker(data *entity.DynamoItem) {

}
