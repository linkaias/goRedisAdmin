package routers

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/routers/view_router"
	"log"
)

func RunApp() {
	router := gin.New()
	router = view_router.RunViewRouter(router)
	err := router.Run(":9527")
	if err != nil {
		log.Fatalln(err)
		return
	}
}
