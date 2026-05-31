package user_router

import (
	"goRedisAdmin/controller/user_controller"

	"github.com/gin-gonic/gin"
)

// UserDataRouter registers authentication-related endpoints.
type UserDataRouter struct{}

// InitRouter mounts user endpoints under /api/v1/user.
func (u UserDataRouter) InitRouter(group *gin.RouterGroup) (R gin.IRoutes) {
	cont := user_controller.NewUserController()
	rou := group.Group("user")
	{
		// Authenticate with configured admin user/password.
		rou.POST("/login", cont.Login)

		// Placeholder logout endpoint.
		rou.GET("/logout", cont.Logout)

	}
	return rou
}
