package callbackprovider

import (
	"strings"

	entity "github.com/sofyan48/cimol/src/entity/http/v1"
)

// InfobipCallback ...
func (callback *ProviderCallback) InfobipCallback(dynamo *entity.DynamoItemResponse,
	data *entity.InfobipCallbackRequest, history *entity.HistoryItem) {
	if validateStatusInfobip(data.Results[0].Status.Name) {
		callback.InfobipCallback(dynamo, data, history)
		return
	}
}

func validateStatusInfobip(status string) bool {
	statusData := []string{
		"UNDELIVERABLE_REJECTED_OPERATOR",
		"UNDELIVERABLE_NOT_DELIVERED",
		"PENDING_ENROUTE",
	}
	for _, i := range statusData {
		if strings.EqualFold(i, status) {
			return true
		}
	}
	return false
}
