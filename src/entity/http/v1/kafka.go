package v1

import "time"

// StateFullFormatKafka ...
type StateFullFormatKafka struct {
	UUID      string            `json:"__uuid" bson:"__uuid"`
	Action    string            `json:"__action" bson:"__action"`
	Data      map[string]string `json:"data" bson:"data"`
	CreatedAt *time.Time        `json:"created_at" bson:"created_at"`
}