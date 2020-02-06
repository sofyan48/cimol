package v1

// DynamoItem ..
type DynamoItem struct {
	ID              string
	Data            string
	History         string
	ReceiverAddress string
	StatusText      string
	Type            string
}

// DynamoItemResponse ..
type DynamoItemResponse struct {
	ID              interface{} `json:"id"`
	Data            interface{} `json:"data"`
	History         interface{} `json:"history"`
	ReceiverAddress interface{} `json:"receiverAddress"`
	StatusText      interface{} `json:"statusText"`
	Type            interface{} `json:"type"`
}
