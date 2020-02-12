package v1

// InfobipRequestPayload ...
type InfobipRequestPayload struct {
	Messages []InfobipMessages `json:"messages"`
}

// // InfobipCallbackRequestPost ...
// type InfobipCallbackRequestPost struct {
// 	Messages []struct {
// 		To     string `json:"to"`
// 		Status struct {
// 			GroupID     uint   `json:"groupId"`
// 			GroupName   string `json:"groupName"`
// 			ID          uint   `json:"id"`
// 			Name        string `json:"name"`
// 			Description string `json:"description"`
// 		} `json:"status"`
// 		MessageID    string `json:"messageId"`
// 		CallbackData string `json:"callbackData"`
// 	} `json:"messages"`
// }

// InfobipMessages .../
type InfobipMessages struct {
	From             string               `json:"from"`
	Destinations     []InfobipDestination `json:"destinations"`
	Text             string               `json:"text"`
	NotifyURL        string               `json:"notifyUrl"`
	NotifyContenType string               `json:"notifyContentType"`
	CallbackData     string               `json:"callbackData"`
}

// InfobipDestination ...
type InfobipDestination struct {
	To string `json:"to"`
}

// InfobipCallbackRequest ...
type InfobipCallbackRequest struct {
	Results []InfobipRequestChild `json:"results"`
}

// InfobipRequestChild ...
type InfobipRequestChild struct {
	BulkID       string             `json:"bulkId"`
	MessagesID   string             `json:"messageId"`
	To           string             `json:"to"`
	SentAt       string             `json:"sentAt"`
	DoneAt       string             `json:"doneAt"`
	SmsCount     uint               `json:"smsCount"`
	MccMnc       string             `json:"mccMnc"`
	Price        PriceChildInfobip  `json:"price"`
	Status       StatusChildInfobip `json:"status"`
	Error        ErrorChildInfobip  `json:"error"`
	CallbackData string             `json:"callbackData"`
}

// PriceChildInfobip ...
type PriceChildInfobip struct {
	PricePerMessages string `json:"pricePerMessages"`
	Currency         string `json:"currency"`
}

// StatusChildInfobip ..
type StatusChildInfobip struct {
	GroupID     uint   `json:"groupId"`
	GroupName   string `json:"groupName"`
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// ErrorChildInfobip ...
type ErrorChildInfobip struct {
	GroupID     uint   `json:"groupId"`
	GroupName   string `json:"groupName"`
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Permanent   bool   `json:"permanent"`
}
