package info_router

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/controller/info_controller"
)

type InfoDataRouter struct{}

func (u InfoDataRouter) InitRouter(group *gin.RouterGroup) (R gin.IRoutes) {
	cont := info_controller.NewInfoController()
	rou := group.Group("info")
	{
		rou.GET("/get_info", cont.GetInfo)

	}
	return rou
}
