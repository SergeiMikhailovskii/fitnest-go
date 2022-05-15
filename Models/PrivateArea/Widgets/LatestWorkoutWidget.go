package Widgets

type LatestWorkoutWidget struct {
	Workouts []Workout `json:"workouts"`
}

type Workout struct {
	Name     string  `json:"name"`
	Calories int     `json:"calories"`
	Minutes  int     `json:"minutes"`
	Progress float32 `json:"progress"`
	Image    string  `json:"image"`
}
