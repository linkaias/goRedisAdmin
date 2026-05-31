package db_router

import (
	"goRedisAdmin/controller/db_data_controller"

	"github.com/gin-gonic/gin"
)

// DbDataRouter registers endpoints for Redis DB/key/value operations.
type DbDataRouter struct{}

// InitRouter mounts all DB data endpoints under /api/v1/db.
func (u DbDataRouter) InitRouter(group *gin.RouterGroup) (R gin.IRoutes) {
	cont := db_data_controller.NewDbDataController()
	rou := group.Group("db")
	{
		// List DB indices and key counts.
		rou.GET("/db_list", cont.DbList)
		// List keys in selected DB with optional filter.
		rou.GET("/get_keys", cont.GetKeys)
		// Get value by key/type payload.
		rou.GET("/get_val", cont.GetVal)
		// Delete one or multiple keys.
		rou.DELETE("/key", cont.DelKey)
		// Export selected keys to JSON.
		rou.POST("/export_keys", cont.ExportKey)
		// Add value with specific Redis data type.
		rou.POST("/key", cont.AddVal)
		// Update key expiration (seconds or persist).
		rou.POST("/key/expire", cont.ExpireKey)
		// Flush current DB or all DBs.
		rou.DELETE("/flush", cont.Flush)
		// Get value by key for specific type-specific view behavior.
		rou.POST("/get_val_by_key", cont.GetValByKey)
	}
	return rou
}
