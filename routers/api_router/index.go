package api_router

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/routers/api_router/db_router"
	"goRedisAdmin/routers/api_router/info_router"
	"goRedisAdmin/routers/api_router/user_router"
	"goRedisAdmin/routers/middleware"
)

var EnterRouter = new(routerGroup)

// 全部路由组
type routerGroup struct {
	db_router.DbRouterGroup
	user_router.UserRouterGroup
	info_router.InfoRouterGroup
}

func RunApiRouter(r *gin.RouterGroup) {
	r.Use(middleware.IpCheckMiddleware())
	gp := r.Group("api/v1")
	EnterRouter.UserRouterGroup.UserDataRouter.InitRouter(gp)

	gp.Use(middleware.AuthorizeJWT())
	EnterRouter.DbRouterGroup.DbDataRouter.InitRouter(gp)
	EnterRouter.InfoDataRouter.InitRouter(gp)

}
