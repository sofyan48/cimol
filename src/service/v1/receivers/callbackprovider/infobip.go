package callbackprovider

import (
	"fmt"

	entity "github.com/sofyan48/rll-daemon-new/src/entity/http/v1"
)

// InfobipCallback ...
func (callback *ProviderCallback) InfobipCallback(dynamo *entity.DynamoItemResponse,
	data *entity.InfobipCallbackRequest) {
	fmt.Println(dynamo)
	fmt.Println(data)
}

func validateStatusInfobip(status *entity.StatusChildInfobip) bool {
	return status.Name != "DELIVERED_TO_HANDSET"
}
