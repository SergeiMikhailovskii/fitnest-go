package Registration

type GoalStepModel struct {
	Goal string `json:"goal" binding:"required"`
}
