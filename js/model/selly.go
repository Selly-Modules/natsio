package jsmodel

// PushNotification ...
type PushNotification struct {
	User        string              `json:"user"`
	Type        string              `json:"type"`
	TargetID    string              `json:"targetId"`
	IsFromAdmin bool                `json:"isFromAdmin"`
	Category    string              `json:"category"`
	Options     NotificationOptions `json:"options"`
}

// NotificationOptions ...
type NotificationOptions struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
