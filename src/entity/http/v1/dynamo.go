package v1

// DynamoItem ..
type DynamoItem struct {
	ID              string `json:"id"`
	Data            string `json:"data"`
	History         string `json:"history"`
	ReceiverAddress string `json:"receiverAddress"`
	StatusText      string `json:"statusText"`
	Type            string `json:"type"`
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
