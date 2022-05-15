package DB

type ActivityAim struct {
	ID                int `json:"id"`
	UserId            int `json:"user_id"`
	WaterIntakeAmount int `json:"water_intake_amount"`
	CaloriesAmount    int `json:"calories_amount"`
}

func (b *ActivityAim) TableName() string {
	return "activity_aim"
}
