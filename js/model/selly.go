package jsmodel

// PushNotification ...
type PushNotification struct {
	User        string `json:"user"`
	Type        string `json:"type"`
	TargetID    string `json:"targetId"`
	IsFromAdmin bool   `json:"isFromAdmin"`
	Category    string `json:"category"`
	Options     struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	} `json:"options"`
}
