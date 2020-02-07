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
