package user_router

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (u UserRouter) InitRouter(group *gin.RouterGroup) (R gin.IRoutes) {
	rou := group.Group("user")
	{
		rou.GET("/login", func(ctx *gin.Context) {

		})
	}
	return rou
}
