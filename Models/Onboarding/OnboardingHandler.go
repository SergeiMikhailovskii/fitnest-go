package Onboarding

import "TestProject/Config"

func CreateDefaultRecord(userId int) error {
	onboardingRecord := Onboarding{
		UserID: userId,
	}
	err := Config.DB.Create(&onboardingRecord).Error

	return err
}

func GetRecordByUserId(userId int, onboarding *Onboarding) error {
	err := Config.DB.Where("user_id = ?", userId).First(onboarding).Error
	return err
}
