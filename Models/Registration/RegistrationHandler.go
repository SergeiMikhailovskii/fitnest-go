package Registration

import (
	"TestProject/Config"
)

func CreatePrimaryRegistrationRecord(userId int) error {
	primaryRecord := PrimaryInfo{
		UserID: userId,
	}
	err := Config.DB.Create(&primaryRecord).Error

	return err
}

func GetPrimaryRegistrationRecordByUserId(userId int, primaryRecord *PrimaryInfo) error {
	err := Config.DB.Where("user_id = ?", userId).First(primaryRecord).Error
	return err
}

func SaveCreateAccountRegistrationRecordByUserId(userId int, model CreateStepModel) error {
	primaryRecord := PrimaryInfo{
		UserID:    userId,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Password:  model.Password,
		Email:     model.Email,
	}
	err := Config.DB.Where("user_id = ?", userId).Updates(primaryRecord).Error
	return err
}
