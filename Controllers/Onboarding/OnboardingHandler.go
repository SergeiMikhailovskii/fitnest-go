package Onboarding

import (
	"TestProject/Models"
	"TestProject/Models/Base"
	"TestProject/Models/Onboarding"
	"github.com/gin-gonic/gin"
)

func getOnboardingStep(c *gin.Context) string {
	cookie, err := c.Cookie(Base.AuthUserCookie.Name)
	if err != nil {
		panic(err)
	}

	user := Models.ParseJwt(cookie)

	onboardingRecord := Onboarding.Onboarding{}
	_ = Onboarding.GetOnboardingRecordByUserId(user.ID, &onboardingRecord)

	if onboardingRecord.ForthStepComplete {
		return "after"
	} else if onboardingRecord.ThirdStepComplete {
		return "step4"
	} else if onboardingRecord.SecondStepComplete {
		return "step3"
	} else if onboardingRecord.FirstStepComplete {
		return "step2"
	} else if !onboardingRecord.FirstStepComplete {
		return "step1"
	} else {
		return "undefined"
	}
}

func submitOnboardingStep(c *gin.Context) string {
	cookie, err := c.Cookie(Base.AuthUserCookie.Name)
	if err != nil {
		panic(err)
	}

	user := Models.ParseJwt(cookie)

	onboardingRecord := Onboarding.Onboarding{}
	_ = Onboarding.GetOnboardingRecordByUserId(user.ID, &onboardingRecord)

	if !onboardingRecord.FirstStepComplete {
		onboardingRecord.FirstStepComplete = true
	} else if !onboardingRecord.SecondStepComplete {
		onboardingRecord.SecondStepComplete = true
	} else if !onboardingRecord.ThirdStepComplete {
		onboardingRecord.ThirdStepComplete = true
	} else if !onboardingRecord.ForthStepComplete {
		onboardingRecord.ForthStepComplete = true
	} else {
		return "error"
	}

	Onboarding.UpdateOnboardingRecord(onboardingRecord)

	return "success"
}
