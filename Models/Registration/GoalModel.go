package Registration

type GoalModel struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Goal   string `json:"goal"`
}

func (b *GoalModel) TableName() string {
	return "goal_model"
}
