package Onboarding

type Onboarding struct {
	ID                 int  `json:"id"`
	UserID             int  `json:"user_id"`
	FirstStepComplete  bool `gorm:"default: false" json:"first_step_complete"`
	SecondStepComplete bool `gorm:"default: false" json:"second_step_complete"`
	ThirdStepComplete  bool `gorm:"default: false" json:"third_step_complete"`
}

func (b *Onboarding) TableName() string {
	return "onboarding"
}
