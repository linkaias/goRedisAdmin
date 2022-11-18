package user_router

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/controller/user_controller"
)

type UserDataRouter struct{}

func (u UserDataRouter) InitRouter(group *gin.RouterGroup) (R gin.IRoutes) {
	cont := user_controller.NewUserController()
	rou := group.Group("user")
	{
		rou.POST("/login", cont.Login)

		rou.GET("/logout", cont.Logout)

	}
	return rou
}
