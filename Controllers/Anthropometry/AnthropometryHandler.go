package Anthropometry

import (
	"TestProject/Config"
	"TestProject/Models/Registration"
	"errors"
	"gorm.io/gorm"
)

func HasAnthropometryRecordByUserId(userId int) bool {
	err := Config.DB.Where("user_id = ?", userId).First(Registration.AnthropometryModel{}).Error
	return errors.Is(err, gorm.ErrRecordNotFound)
}
