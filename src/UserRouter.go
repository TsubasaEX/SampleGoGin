package src

import (
	"SampleGoGin/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	user := r.Group("/users")
	user.GET("/", service.FindAllUsers)
	user.POST("/", service.PostUser)
	user.DELETE("/:id", service.DeleteUser)
	user.PUT("/:id", service.PutUser)
}
