package pojo

import (
	"SampleGoGin/database"
	"log"

	"gopkg.in/mgo.v2/bson"
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
	return result.RowsAffected > 0
}

func UpdateUser(userId int, user User) User {
	newUser := User{}
	database.DBConnect.Model(&newUser).Where("id = ?", userId).Updates(user)
	return newUser
}

func CheckUserPassword(name string, password string) User {
	user := User{}
	database.DBConnect.Where("name = ? and password = ?", name, password).First(&user)
	return user
}

// MongoDB
// -------------------------------------------------------------
func MgoCreateUser(user User) User {
	database.MgoConnect.Insert(user)
	return user
}

func MgoFindAllUsers() []User {
	var users []User
	database.MgoConnect.Find(nil).All(&users)
	return users
}

func MgoFindById(userId int) User {
	user := User{}
	database.MgoConnect.Find(bson.M{
		"id": userId,
	}).One(&user)
	return user
}

func MgoUpdateUser(userId int, user User) User {
	updateUserId := bson.M{"id": userId}
	updateData := bson.M{"$set": user}
	err := database.MgoConnect.Update(updateUserId, updateData)
	if err != nil {
		log.Println("Error :" + err.Error())
		return User{}
	}
	return user
}

// DeleteUser
func MgoDeleteUser(userId int) bool {
	err := database.MgoConnect.Remove(bson.M{
		"id": userId,
	})
	if err != nil {
		log.Println("Error :" + err.Error())
		return false
	}
	return true
}
