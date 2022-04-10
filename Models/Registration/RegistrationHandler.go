package Registration

import (
	"TestProject/Config"
	"TestProject/Util"
	"time"
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

func SaveCompleteAccountRegistrationRecordByUserId(userId int, model CompleteStepModel) error {
	date, err := time.Parse(Util.DD_MM_YYYY, model.DateOfBirth)
	if err != nil {
		return err
	}
	primaryRecord := PrimaryInfo{
		UserID:    userId,
		Sex:       model.Sex,
		BirthDate: date,
	}
	anthropometryRecord := AnthropometryModel{
		UserID: userId,
		Height: model.Height,
		Weight: model.Weight,
	}
	err = Config.DB.Where("user_id = ?", userId).Updates(primaryRecord).Error
	if err != nil {
		return err
	}

	err = Config.DB.Create(&anthropometryRecord).Error
	return err
}

func SaveGoalRegistrationRecordByUserId(userId int, model GoalStepModel) error {
	goalRecord := GoalModel{
		UserID: userId,
		Goal:   model.Goal,
	}

	err := Config.DB.Create(&goalRecord).Error
	return err
}

func SaveWelcomeBackRegistrationRecordByUserId(userId int) error {
	goalRecord := PrimaryInfo{
		UserID:            userId,
		WelcomeBackSubmit: true,
	}

	err := Config.DB.Create(&goalRecord).Error
	return err
}
