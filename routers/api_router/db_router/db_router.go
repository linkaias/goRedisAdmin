package db_router

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/controller/db_data_controller"
)

type DbDataRouter struct{}

func (u DbDataRouter) InitRouter(group *gin.RouterGroup) (R gin.IRoutes) {
	cont := db_data_controller.NewDbDataController()
	rou := group.Group("db")
	{
		//数据库列表
		rou.GET("/db_list", cont.DbList)
		rou.GET("/get_keys", cont.GetKeys)

	}
	return rou
}
