package Registration

type AnthropometryModel struct {
	ID     int     `json:"id"`
	UserID int     `json:"user_id"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}

func (b *AnthropometryModel) TableName() string {
	return "anthropometry_model"
}
