package Util

import "errors"

var UserNotFound = errors.New("error.userNotFound")
var UserExists = errors.New("error.userExists")