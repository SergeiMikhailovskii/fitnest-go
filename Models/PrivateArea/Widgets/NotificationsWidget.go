package Widgets

type NotificationsWidget struct {
	Notifications []Notification `json:"notifications"`
}

type Notification struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Date     string `json:"date"`
	Type     string `json:"type"`
	IsActive bool   `json:"is_active"`
	IsPinned bool   `json:"is_pinned"`
}
