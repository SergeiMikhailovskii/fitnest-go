package Goal

import (
	"TestProject/Config"
	"TestProject/Models/Registration"
	"errors"
	"gorm.io/gorm"
)

func HasGoalRecordByUserId(userId int) bool {
	err := Config.DB.Where("user_id = ?", userId).First(Registration.GoalModel{}).Error
	return errors.Is(err, gorm.ErrRecordNotFound)
}
