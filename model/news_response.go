package model

// GetProductNoticesByInventoryResponse ....
type GetProductNoticesByInventoryResponse struct {
	Notices []NewsAppResponse `json:"notices"`
}

// NewsAppResponse ...
type NewsAppResponse struct {
	ID           string        `json:"_id"`
	Title        string        `json:"title,omitempty"`
	Target       *TargetNewDoc `json:"target,omitempty"`
	ActionType   *ActionType   `json:"action"`
	ShortDesc    string        `json:"shortDesc,omitempty"`
	Type         string        `json:"type"`
	ShortTitle   string        `json:"shortTitle,omitempty"`
	Color        string        `json:"color"`
	Options      *NewsOptions  `json:"options,omitempty"`
	DisplayStyle string        `json:"displayStyle"`
}

// NewsOptions ...
type NewsOptions struct {
	Category string `json:"category"`
}

// TargetNewDoc ...
type TargetNewDoc struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

// ActionType ...
type ActionType struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Text  string `json:"text,omitempty"`
}
