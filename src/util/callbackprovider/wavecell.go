package callbackprovider

import (
	"strings"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

// WavecellCallback ...
func (callback *ProviderCallback) WavecellCallback(dynamo *entity.DynamoItemResponse,
	data *entity.WavecellCallBackRequest, history *entity.HistoryItem) {
	if validateStatusWavecell(data.Status) {
		callback.wavecellMessagesNotSuccess(dynamo, data, history)
		return
	}
}

func validateStatusWavecell(status string) bool {
	statusData := []string{
		"REJECTED BY CARRIER",
		"REJECTED BY DEVICE",
		"REJECTED BY WAVECELL",
	}
	for _, i := range statusData {
		if strings.EqualFold(i, status) {
			return true
		}
	}
	return false
}
