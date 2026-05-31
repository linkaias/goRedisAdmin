package info_router

import (
	"goRedisAdmin/controller/info_controller"

	"github.com/gin-gonic/gin"
)

// InfoDataRouter registers endpoints for Redis INFO inspection.
type InfoDataRouter struct{}

// InitRouter mounts info endpoints under /api/v1/info.
func (u InfoDataRouter) InitRouter(group *gin.RouterGroup) (R gin.IRoutes) {
	cont := info_controller.NewInfoController()
	rou := group.Group("info")
	{
		// Returns Redis INFO output formatted for HTML rendering.
		rou.GET("/get_info", cont.GetInfo)

	}
	return rou
}
