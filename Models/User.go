package Models

import (
	"TestProject/Config"
	"TestProject/Errors"
	"gorm.io/gorm"
)

func CreateUser(User *User) (err error) {
	err = Config.DB.Create(&User).Error
	return err
}

func GetUsers(db *gorm.DB, Users *[]User) (err error) {
	err = db.Find(Users).Error
	return err
}

func GetUser(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	return err
}

func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id = ?", id).Delete(User).Error
	return err
}

func UpdateUser(db *gorm.DB, User *User) (err error) {
	err = db.Save(User).Error
	return err
}

func GetUserByCreds(user *User, userRequest User) (err error) {
	if err = Config.DB.Where("Login = ? AND Password = ?", userRequest.Login, userRequest.Password).First(user).Error; err != nil {
		err = Errors.UserNotFound
	}
	return err
}

func CheckUserExistsByLogin(userRequest User) (err error) {
	user := User{}
	err = Config.DB.Where("Login = ?", userRequest.Login).First(&user).Error
	if err == nil {
		err = Errors.UserExists
	} else {
		err = Errors.UserNotFound
	}
	return err
}
