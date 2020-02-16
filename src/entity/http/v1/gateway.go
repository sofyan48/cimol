package v1

// PostNotificationResponse ..
type PostNotificationResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}

// PayloadPostNotificationRequest ...
type PayloadPostNotificationRequest struct {
	OTP    bool   `json:"otp"`
	Msisdn string `json:"msisdn"`
	Text   string `json:"text"`
}

// PostNotificationRequest ...
type PostNotificationRequest struct {
	Type    string                          `json:"type"`
	UUID    string                          `json:"uuid"`
	Payload *PayloadPostNotificationRequest `json:"payload"`
}

// PostNotificationRequestEmail ...
type PostNotificationRequestEmail struct {
	Type    string               `json:"type"`
	UUID    string               `json:"uuid"`
	Payload *PayloadRequestEmail `json:"payload"`
}

// PayloadRequestEmail ...
type PayloadRequestEmail struct {
	To         string            `json:"to"`
	From       string            `json:"from"`
	Subject    string            `json:"subject"`
	TemplateID string            `json:"template_id"`
	Data       map[string]string `json:"data"`
}

// PostNotificationRequestPush ...
type PostNotificationRequestPush struct {
	Type string `json:"type"`
	UUID string `json:"uuid"`
}
