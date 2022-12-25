package Widgets

type ProfileInfoWidget struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Program   string `json:"program"`
	Height    int    `json:"height"`
	Weight    int    `json:"weight"`
	Age       int    `json:"age"`
}

type ProfilePrimaryInfoQuery struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	DatePart  int    `json:"date_part"`
}
