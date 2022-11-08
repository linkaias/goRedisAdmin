package routers

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/global/initData"
	"goRedisAdmin/routers/api_router"
	"goRedisAdmin/routers/middleware"
	"goRedisAdmin/routers/view_router"
	"log"
)

func RunApp() {
	router := gin.New()

	//访问图片资源使用，可删除
	router.Static("/image", "./image")

	//登录限制
	router.Use(middleware.HTTPAuthMiddleware())
	//渲染视图
	router = view_router.RunViewRouter(router)
	//API Router
	api_router.RunApiRouter(router.Group(""))

	err := router.Run(getRunPort())
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func getRunPort() string {
	cfg := initData.IniRead.Section("admin")
	return ":" + cfg.Key("port").String()
}
