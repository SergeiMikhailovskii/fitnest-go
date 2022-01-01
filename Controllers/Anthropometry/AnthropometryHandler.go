package Anthropometry

import (
	"TestProject/Config"
	"TestProject/Models/Registration"
)

func HasAnthropometryRecordByUserId(userId int) (bool, error) {
	sample := Config.DB.Where("user_id = ?", userId).First(Registration.AnthropometryModel{})
	return sample.RowsAffected != 0, sample.Error
}
