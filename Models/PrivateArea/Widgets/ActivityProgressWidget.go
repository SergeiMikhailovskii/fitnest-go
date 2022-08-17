package Widgets

type ActivityProgressWidget struct {
	Progresses []ActivityProgressItem `json:"progresses"`
}

type ActivityProgressItem struct {
	Date  string
	Total int
}
