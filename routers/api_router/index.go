package api_router

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/routers/api_router/db_router"
)

var EnterRouter = new(routerGroup)

//全部路由组
type routerGroup struct {
	db_router.DbRouterGroup
}

func RunApiRouter(r *gin.RouterGroup) {
	gp := r.Group("api/v1")
	EnterRouter.DbRouterGroup.DbDataRouter.InitRouter(gp)
}
