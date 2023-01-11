package pojo

import (
	"SampleGoGin/database"
	"log"
)

type User struct {
	Id       int    `json:"UserId"`
	Name     string `json:"UserName"`
	Password string `json:"UserPassword"`
	Email    string `json:"UserEmail"`
}

func FindAllUsers() []User {
	var users []User
	database.DBConnect.Find(&users)
	return users
}

func FindByUserId(userId int) User {
	var user User
	err := database.DBConnect.Where("id = ?", userId).Last(&user).Error
	if err != nil { // if not found, the err wiil be “record not found”
		log.Println("Error :" + err.Error())
	}
	return user
}
