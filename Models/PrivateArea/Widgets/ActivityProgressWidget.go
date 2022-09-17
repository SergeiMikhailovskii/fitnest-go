package Widgets

type ActivityProgressWidget struct {
	Progresses []ActivityProgressItem `json:"progresses"`
}

type ActivityProgressItem struct {
	Date  string `json:"date"`
	Total int    `json:"total"`
}
