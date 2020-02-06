package v1

// WavecelllCallBackRequest ...
type WavecelllCallBackRequest struct {
	UmID            string `json:"umid"`
	Timestamp       string `json:"timestamp"`
	Status          string `json:"status"`
	StatusCode      string `json:"statusCode"`
	Error           string `json:"error"`
	ErrorCode       uint   `json:"errorCode"`
	Source          string `json:"Source"`
	SubAccountID    string `json:"subAccountId"`
	Version         string `json:"version"`
	Destination     string `json:"destination"`
	BatchID         string `json:"batchId"`
	ClientMessageID string `json:"clientMessageId"`
	ClientBatchID   string `json:"ClientBatchId"`
	Price           struct {
		Total    float32 `json:"total"`
		PerSMS   float32 `json:"perSms"`
		Currency string  `json:"currency"`
	} `json:"price"`
	SmsCount uint `json:"smsCount"`
}
