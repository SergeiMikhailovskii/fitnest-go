package Goal

import (
	"TestProject/Config"
	"TestProject/Models/Registration"
)

func HasGoalRecordByUserId(userId int) bool {
	err := Config.DB.Where("user_id = ?", userId).First(&Registration.GoalModel{}).Error
	return err == nil
}
