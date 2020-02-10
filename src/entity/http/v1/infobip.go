package v1

// InfobipRequestPayload ...
type InfobipRequestPayload struct {
	Messages []InfobipMessages `json:"messages"`
}

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

// InfobipCallBackRequest ...
type InfobipCallBackRequest struct {
	Results []InfobipRequestChild
}

// InfobipRequestChild ...
type InfobipRequestChild struct {
	BulkID     string
	MessagesID string
	To         string
	SentAt     string
	DoneAt     string
	SmsCount   uint
	MccMnc     string
	Price      struct {
		PricePerMessages string
		Currency         string
	}
	Status struct {
		GroupID     string
		GroupName   string
		ID          string
		Name        string
		Description string
	}
	Error struct {
		GroupID     string
		GroupName   string
		ID          string
		Name        string
		Description string
		Permanent   bool
	}
	CallBackData string
}
