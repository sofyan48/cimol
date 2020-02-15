package v1

// VersionTemplates ..
type VersionTemplates struct {
	ID                   string `json:"id"`
	UserID               uint   `json:"user_id"`
	TemplateID           string `json:"template_id"`
	Active               uint   `json:"active"`
	Name                 string `json:"name"`
	HTMLContent          string `json:"html_content"`
	PlainContent         string `json:"plain_content"`
	GeneratePlainContent bool   `json:"generate_plain_content"`
	Subject              string `json:"subject"`
	UpdateAt             string `json:"updated_at"`
	Editor               string `json:"editor"`
	ThumbnailURI         string `json:"thumbnail_url"`
}

// TemplateResponse ..
type TemplateResponse struct {
	ID         string             `json:"id"`
	Name       string             `json:"name"`
	Generation string             `json:"generation"`
	UpdateAt   string             `json:"updated_at"`
	Versions   []VersionTemplates `json:"versions"`
}

// SendPayload ...
type SendPayload struct {
	Personalization []PersonalizationData
	From            []SenderFrom `json:"from"`
	TemplateID      string       `json:"template_id"`
}

// SenderTo ..
type SenderTo struct {
	Email string `json:"email"`
}

// SenderFrom ...
type SenderFrom struct {
	Email string `json:"email"`
}

// PersonalizationData ...
type PersonalizationData struct {
	Subject       string            `json:"subject"`
	To            []SenderTo        `json:"to"`
	Substitutions map[string]string `json:"substitutions"`
}

// ContentData ...
type ContentData struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
