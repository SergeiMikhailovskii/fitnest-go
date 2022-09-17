package Widgets

type LatestActivityWidget struct {
	Activities []Activity `json:"activities"`
}

type Activity struct {
	ID     int    `json:"id"`
	Amount int    `json:"amount"`
	Type   string `json:"type"`
	Time   string `json:"time"`
}
