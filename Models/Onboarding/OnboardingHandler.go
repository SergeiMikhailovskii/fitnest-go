package Onboarding

import (
	"TestProject/Config"
)

func CreateOnboardingDefaultRecord(userId int) error {
	onboardingRecord := Onboarding{
		UserID: userId,
	}
	err := Config.DB.Create(&onboardingRecord).Error

	return err
}

func GetOnboardingRecordByUserId(userId int, onboarding *Onboarding) error {
	err := Config.DB.Where("user_id = ?", userId).First(onboarding).Error
	return err
}

func UpdateOnboardingRecord(onboarding Onboarding) error {
	err := Config.DB.Save(onboarding).Error
	return err
}
