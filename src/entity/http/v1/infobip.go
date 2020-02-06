package v1

// CallBackRequest ...
type InfobipCallBackRequest struct {
	Messages []struct {
		To     string `json:"to"`
		Status struct {
			GroupID     uint   `json:"groupId"`
			GroupName   string `json:"groupName"`
			ID          uint   `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"status"`
		MessageID string `json:"messageId"`
	} `json:"messages"`
}

// InfobipResponse ...
type InfobipResponse struct {
	Results []InfobipResponseChild
}

// InfobipResponseChild ...
type InfobipResponseChild struct {
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
