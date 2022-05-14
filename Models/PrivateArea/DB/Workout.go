package DB

type Workout struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Calories int    `json:"calories"`
	Minutes  int    `json:"minutes"`
}

func (b *Workout) TableName() string {
	return "workout"
}
