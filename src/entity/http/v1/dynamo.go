package v1

import "time"

// DynamoItem ..
type DynamoItem struct {
	ID              string `json:"id"`
	Data            string `json:"data"`
	History         string `json:"history"`
	ReceiverAddress string `json:"receiverAddress"`
	StatusText      string `json:"statusText"`
	Type            string `json:"type"`
}
