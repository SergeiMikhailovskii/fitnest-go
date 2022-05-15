package DB

type UserWorkout struct {
	ID        int     `json:"id"`
	UserId    int     `json:"user_id"`
	WorkoutId int     `json:"workout_id"`
	Progress  float32 `json:"progress"`
}

func (b *UserWorkout) TableName() string {
	return "user_workout"
}
