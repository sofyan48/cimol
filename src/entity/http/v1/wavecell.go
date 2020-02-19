package v1

// WavecelllCallBackRequest ...
type WavecelllCallBackRequest struct {
	UmID            string `json:"umid"`
	Timestamp       string `json:"timestamp"`
	Status          string `json:"status"`
	StatusCode      uint   `json:"statusCode"`
	Error           string `json:"error"`
	ErrorCode       uint   `json:"errorCode"`
	Source          string `json:"Source"`
	SubAccountID    string `json:"subAccountId"`
	Version         uint   `json:"version"`
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

// WavecellRequest ...
type WavecellRequest struct {
	Source          string `json:"source"`
	Destination     string `json:"destination"`
	Text            string `json:"text"`
	ClientMessageID string `json:"clientMessageId"`
	DLRCallback     string `json:"dlrCallbackUrl"`
}

// WavecellCallbackRequest ...
type WavecellCallbackRequest struct {
	Messages []WavecellCallbackMessages `json:"messages"`
}

// WavecellCallbackMessages ...
type WavecellCallbackMessages struct {
	To        string                   `json:"to"`
	Status    []WavecellResponseStatus `json:"status"`
	MessageID string                   `json:"messageId"`
}

// WavecellCallbackStatus ...
type WavecellCallbackStatus struct {
	GroupID     uint   `json:"groupId"`
	GroupName   string `json:"groupName"`
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// WavecellResponse ...
type WavecellResponse struct {
	UMID             string                  `json:"umid"`
	ClientMessagesID string                  `json:"clientMessageId"`
	Destination      string                  `json:"destination"`
	Encoding         string                  `json:"encoding"`
	Status           *WavecellResponseStatus `json:"status"`
}

// WavecellResponseStatus ...
type WavecellResponseStatus struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
