package db_router

import "github.com/gin-gonic/gin"

type DbDataRouter struct{}

func (u DbDataRouter) InitRouter(group *gin.RouterGroup) (R gin.IRoutes) {
	rou := group.Group("db")
	{
		rou.GET("/list", func(ctx *gin.Context) {

		})
	}
	return rou
}
