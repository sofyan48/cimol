package v1

// DynamoItem ..
type DynamoItem struct {
	ID              string                  `json:"id"`
	Data            string                  `json:"data"`
	History         map[string]*HistoryItem `json:"history"`
	ReceiverAddress string                  `json:"receiverAddress"`
	StatusText      string                  `json:"statusText"`
	Type            string                  `json:"type"`
}

// DynamoItemResponse ..
type DynamoItemResponse struct {
	ID   string `json:"id"`
	Data string `json:"data"`
	// History         interface{} `json:"history"`
	History         []string `json:"history"`
	ReceiverAddress string   `json:"receiverAddress"`
	StatusText      string   `json:"statusText"`
	Type            string   `json:"type"`
}

// HistoryItem ...
type HistoryItem struct {
	Provider       string                          `json:"provider"`
	DeliveryReport string                          `json:"delivery_report"`
	Response       string                          `json:"response"`
	CallbackData   string                          `json:"callback_data"`
	Payload        *PayloadPostNotificationRequest `json:"payload"`
}
