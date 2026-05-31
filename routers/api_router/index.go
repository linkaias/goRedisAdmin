package api_router

import (
	"goRedisAdmin/routers/api_router/db_router"
	"goRedisAdmin/routers/api_router/info_router"
	"goRedisAdmin/routers/api_router/user_router"
	"goRedisAdmin/routers/middleware"

	"github.com/gin-gonic/gin"
)

// EnterRouter is the aggregated router group entry for all API modules.
var EnterRouter = new(routerGroup)

// routerGroup bundles all API route module groups.
type routerGroup struct {
	db_router.DbRouterGroup
	user_router.UserRouterGroup
	info_router.InfoRouterGroup
}

// RunApiRouter mounts all API routes and middlewares.
//
// Order:
//  1. IP whitelist check for all /api/v1 routes
//  2. user login/logout routes (no JWT required)
//  3. JWT middleware for remaining protected routes
//  4. DB and info routes
func RunApiRouter(r *gin.RouterGroup) {
	r.Use(middleware.IpCheckMiddleware())
	gp := r.Group("api/v1")
	EnterRouter.UserRouterGroup.UserDataRouter.InitRouter(gp)

	gp.Use(middleware.AuthorizeJWT())
	EnterRouter.DbRouterGroup.DbDataRouter.InitRouter(gp)
	EnterRouter.InfoDataRouter.InitRouter(gp)

}
