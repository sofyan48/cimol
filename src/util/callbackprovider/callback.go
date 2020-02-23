package callbackprovider

import (
	entity "github.com/sofyan48/cimol/src/entity/http/v1"
	"github.com/sofyan48/cimol/src/util/helper/libaws"
	"github.com/sofyan48/cimol/src/util/logging"
	"github.com/sofyan48/cimol/src/util/provider"
)

// ProviderCallback ...
type ProviderCallback struct {
	AwsLib   libaws.AwsInterface
	Provider provider.ProvidersInterface
	Logs     logging.LogInterface
}

// ProviderCallbackHandler ...
func ProviderCallbackHandler() *ProviderCallback {
	return &ProviderCallback{
		AwsLib:   libaws.AwsHAndler(),
		Provider: provider.ProvidersHandler(),
		Logs:     logging.LogHandler(),
	}
}

// ProviderCallbackInterface ...
type ProviderCallbackInterface interface {
	InfobipCallback(dynamo *entity.DynamoItemResponse, data *entity.InfobipCallbackRequest, history *entity.HistoryItem)
	WavecellCallback(dynamo *entity.DynamoItemResponse, data *entity.WavecellCallBackRequest, history *entity.HistoryItem)
	wavecellMessagesNotSuccess(dynamo *entity.DynamoItemResponse, data *entity.WavecellCallBackRequest, history *entity.HistoryItem)
	infobipMessagesNotSuccess(dynamo *entity.DynamoItemResponse, data *entity.WavecellCallBackRequest, history *entity.HistoryItem)
}
