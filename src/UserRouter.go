package src

import (
	session "SampleGoGin/middlewares"
	"SampleGoGin/pojo"
	"SampleGoGin/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) {
	// user := r.Group("/users")
	// or you could set the middleware(session.SetSession()) in router
	user := r.Group("/users", session.SetSession())

	// user.GET("/", service.FindAllUsers)
	user.GET("/", service.CacheUserAllDecorator(service.RedisUserAll,
		"user_all", []pojo.User{}))
	// user.GET("/:id", service.FindByUserId)
	user.GET("/:id", service.CacheOneUserDecorator(service.RedisOneUser,
		"id", "user_%s", pojo.User{}))
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
