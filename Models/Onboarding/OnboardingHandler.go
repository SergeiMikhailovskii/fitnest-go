package Onboarding

import "TestProject/Config"

func CreateDefaultRecord(userId int) error {
	onboardingRecord := Onboarding{
		UserID: userId,
	}
	err := Config.DB.Create(&onboardingRecord).Error

	return err
}
