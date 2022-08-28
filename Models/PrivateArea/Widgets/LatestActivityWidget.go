package Widgets

type LatestActivityWidget struct {
	Activities []Activity `json:"activities"`
}

type Activity struct {
	Amount int
	Type   string
	Time   string
}
