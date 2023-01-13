package service

import (
	"SampleGoGin/pojo"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// var userList = []pojo.User{}

func FindAllUsers(c *gin.Context) {
	// c.JSON(http.StatusOK, userList)
	users := pojo.FindAllUsers()
	c.JSON(http.StatusOK, users)
}

func FindByUserId(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	user := pojo.FindByUserId(userId)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}
	log.Println("User ->", user)
	c.JSON(http.StatusOK, user)
}

func PostUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error : "+err.Error())
		return
	}
	// userList = append(userList, user)
	newUser, err := pojo.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error : "+err.Error())
		return
	}
	c.JSON(http.StatusOK, newUser)
}

func DeleteUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if !pojo.DeleteUser(userId) {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}
	c.JSON(http.StatusOK, "Successfully deleted")
	// for index, user := range userList {
	// 	if user.Id == userId {
	// 		userList = append(userList[:index], userList[index+1:]...)
	// 		c.JSON(http.StatusOK, "Successfully deleted")
	// 		return
	// 	}
	// }
	// c.JSON(http.StatusNotFound, "Not Found")
}

func PutUser(c *gin.Context) {
	user := pojo.User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error : "+err.Error())
	}
	userId, _ := strconv.Atoi(c.Param("id"))
	user = pojo.UpdateUser(userId, user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Not Found")
		return
	}
	c.JSON(http.StatusOK, user)

	// userId, _ := strconv.Atoi(c.Param("id"))
	// for index, user := range userList {
	// 	if user.Id == userId {
	// 		userList[index] = beforeUser
	// 		log.Println(userList[index])
	// 		c.JSON(http.StatusOK, "Successfully put")
	// 		return
	// 	}
	// }
	// c.JSON(http.StatusNotFound, "Not Found")
}
