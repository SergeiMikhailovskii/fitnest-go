package Widgets

import "time"

type ActivityStatusWidget struct {
	HeartRate   *HeartRateSubWidget   `json:"heart_rate"`
	WaterIntake *WaterIntakeSubWidget `json:"water_intake"`
	Sleep       *SleepSubWidget       `json:"sleep"`
	Calories    CaloriesSubWidget     `json:"calories"`
}

type HeartRateSubWidget struct {
	Rate int       `json:"rate"`
	Date time.Time `json:"date"`
}

type WaterIntakeSubWidget struct {
	Amount   int           `json:"amount"`
	Progress float32       `json:"progress"`
	Intakes  []WaterIntake `json:"intakes"`
}

type WaterIntake struct {
	TimeDiapason   string `json:"time_diapason"`
	AmountInMillis int    `json:"amount_in_millis"`
}

type SleepSubWidget struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}

type CaloriesSubWidget struct {
	Consumed int `json:"consumed"`
	Left     int `json:"left"`
}
