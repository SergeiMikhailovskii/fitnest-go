package Util

import "errors"

var UserNotFound = errors.New("error.userNotFound")
var UserExists = errors.New("error.userExists")

var OnboardingFinished = errors.New("onboarding.finished")
var OnboardingStepNotFound = errors.New("onboarding.step_not_found")
