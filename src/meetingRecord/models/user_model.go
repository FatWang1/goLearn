package models

import (
	"github.com/jinzhu/gorm"
	"meetingRecord/utils"
)

type Usr struct {
	Username string
	Password string
	Status   int
}

type User struct {
	gorm.Model
	Usr
}

func QueryUserRowForLogin(username string, md5Password string) (usr User) {
	utils.DB.Where("username = ? AND password = ?", username, md5Password).Find(&usr)
	return usr
}

func QueryUserExist(username string) bool {
	var count int
	utils.DB.Model(User{}).Where("username = ?", username).Count(&count)
	if count != 0 {
		return true
	}
	return false
}

func InsertUser(user User) {
	utils.DB.Create(&user)
}
