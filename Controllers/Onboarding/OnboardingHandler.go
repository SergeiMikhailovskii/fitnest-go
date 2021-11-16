package Onboarding

import (
	"TestProject/Models"
	"TestProject/Models/Base"
	"TestProject/Models/Onboarding"
	"TestProject/Util"
	"github.com/gin-gonic/gin"
)

func IsOnboardingFinished(c *gin.Context) bool {
	onboardingRecord := getOnboardingRecord(c)
	return onboardingRecord.FirstStepComplete &&
		onboardingRecord.SecondStepComplete &&
		onboardingRecord.ThirdStepComplete &&
		onboardingRecord.ForthStepComplete
}

func getOnboardingStep(c *gin.Context) string {
	onboardingRecord := getOnboardingRecord(c)
	if onboardingRecord.ForthStepComplete {
		return "STEP_AFTER_ONBOARDING"
	} else if onboardingRecord.ThirdStepComplete {
		return "STEP_FORTH_ONBOARDING"
	} else if onboardingRecord.SecondStepComplete {
		return "STEP_THIRD_ONBOARDING"
	} else if onboardingRecord.FirstStepComplete {
		return "STEP_SECOND_ONBOARDING"
	} else if !onboardingRecord.FirstStepComplete {
		return "STEP_FIRST_ONBOARDING"
	} else {
		return "STEP_UNDEFINED_ONBOARDING"
	}
}

func submitOnboardingStep(c *gin.Context) error {
	onboardingRecord := getOnboardingRecord(c)

	if IsOnboardingFinished(c) {
		return Util.OnboardingFinished
	} else if !onboardingRecord.FirstStepComplete {
		onboardingRecord.FirstStepComplete = true
	} else if !onboardingRecord.SecondStepComplete {
		onboardingRecord.SecondStepComplete = true
	} else if !onboardingRecord.ThirdStepComplete {
		onboardingRecord.ThirdStepComplete = true
	} else if !onboardingRecord.ForthStepComplete {
		onboardingRecord.ForthStepComplete = true
	}

	return Onboarding.UpdateOnboardingRecord(onboardingRecord)
}

func getOnboardingRecord(c *gin.Context) Onboarding.Onboarding {
	cookie, err := c.Cookie(Base.AuthUserCookie.Name)
	if err != nil {
		panic(err)
	}

	user := Models.ParseJwt(cookie)
	onboardingRecord := Onboarding.Onboarding{}
	_ = Onboarding.GetOnboardingRecordByUserId(user.ID, &onboardingRecord)
	return onboardingRecord
}
