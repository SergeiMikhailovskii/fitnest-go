package DB

type WaterIntakeAim struct {
	ID     int `json:"id"`
	UserId int `json:"user_id"`
	Amount int `json:"amount"`
}

func (b *WaterIntakeAim) TableName() string {
	return "water_intake_aim"
}
