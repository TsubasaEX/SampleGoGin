package src

import (
	session "SampleGoGin/middlewares"
	"SampleGoGin/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	// user := r.Group("/users")
	// or you could set the middleware(session.SetSession()) in router
	user := r.Group("/users", session.SetSession())

	user.GET("/", service.FindAllUsers)
	user.GET("/:id", service.FindByUserId)
	user.POST("/", service.PostUser)
	user.POST("/more", service.CreateUserList)

	user.PUT("/:id", service.PutUser)

	user.POST("/login", service.LoginUser)

	user.GET("/check", service.CheckUserSession)

	//The APIs after this function will require login
	user.Use(session.AuthSession())
	{
		user.DELETE("/:id", service.DeleteUser)
		user.GET("/logout", service.LogoutUser)
	}

}
