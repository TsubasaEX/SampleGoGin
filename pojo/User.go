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
	err := database.DBConnect.Where("id = ?", userId).First(&user).Error
	if err != nil { // if not found, the err wiil be “record not found”
		log.Println("Error :" + err.Error())
	}
	return user
}

func CreateUser(user User) User {
	database.DBConnect.Create(&user)
	return user
}

func DeleteUser(userId int) bool {
	result := database.DBConnect.Where("id = ?", userId).Delete(&User{})
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func UpdateUser(userId int, user User) User {
	newUser := User{}
	database.DBConnect.Model(&newUser).Where("id = ?", userId).Updates(user)
	return newUser
}
