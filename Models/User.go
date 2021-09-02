package Models

import (
	"TestProject/Config"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
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

func GetUserByID(user *User, login string) (err error) {
	if err = Config.DB.Where("Login = ?", login).First(user).Error; err != nil {
		return err
	}
	return nil
}
