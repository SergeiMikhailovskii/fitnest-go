package Anthropometry

import (
	"TestProject/Config"
	"TestProject/Models/Registration"
)

func HasAnthropometryRecordByUserId(userId int) bool {
	err := Config.DB.Where("user_id = ?", userId).First(&Registration.AnthropometryModel{}).Error
	return err == nil
}
