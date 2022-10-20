package api_router

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/routers/api_router/user_router"
)

var EnterRouter = new(routerGroup)

//全部路由组
type routerGroup struct {
	user_router.UserRouterGroup
}

func RunApiRouter(r *gin.RouterGroup) {
	gp := r.Group("api/v1")
	EnterRouter.UserRouterGroup.UserRouter.InitRouter(gp)
}
