package pojo

import (
	"SampleGoGin/database"
	"log"
)

type User struct { //DB : Users
	Id int `json:"UserId" gorm:"primaryKey" binding:"required"` // Id DB: id, UserId DB: user_id
	// Name     string `json:"UserName" gorm:"Column:username"` // Name DB: name, UserName DB:user_name
	Name     string `json:"UserName" binding:"required,gt=5"`
	Password string `json:"UserPassword" binding:"min=4,max=20,userpwd"`
	Email    string `json:"UserEmail" binding:"email"`
}

type Users struct {
	UserList     []User `json:"UserList" binding:"required,gt=0,lt=3"`
	UserListSize int    `json:"UserListSize"`
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

func CreateUser(user User) (User, error) {
	err := database.DBConnect.Create(&user).Error
	if err != nil {
		log.Println("Error :" + err.Error())
		return User{}, err
	}
	return user, nil
}

// func CreateUsers(users []User) error {
func CreateUsers(users ...User) error {
	err := database.DBConnect.Create(users).Error
	if err != nil {
		log.Println("Error :" + err.Error())
		return err
	}
	return nil
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
