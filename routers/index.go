package routers

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/routers/api_router"
	"goRedisAdmin/routers/view_router"
	"log"
)

func RunApp() {
	router := gin.New()
	//渲染视图
	router = view_router.RunViewRouter(router)
	//API
	api_router.RunApiRouter(router.Group(""))

	err := router.Run(":9527")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
