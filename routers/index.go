package routers

import (
	"goRedisAdmin/global/initData"
	"goRedisAdmin/routers/api_router"
	"goRedisAdmin/routers/view_router"
	"log"

	"github.com/gin-gonic/gin"
)

// RunApp constructs the Gin engine, mounts static/view/api routers,
// and starts the HTTP server.
func RunApp() {
	router := gin.New()

	// Static image route used by the current UI assets.
	router.Static("/image", "./image")

	// Optional basic-auth middleware (disabled by default).
	//router.Use(middleware.HTTPAuthMiddleware())
	// Register static frontend page routes.
	router = view_router.RunViewRouter(router)
	// Register API routes under /api/v1.

	api_router.RunApiRouter(router.Group(""))

	err := router.Run(getRunPort())
	if err != nil {
		log.Fatalln(err)
		return
	}
}

// getRunPort reads backend port from [admin] section in config.ini.
func getRunPort() string {
	cfg := initData.IniRead.Section("admin")
	return ":" + cfg.Key("port").String()
}
