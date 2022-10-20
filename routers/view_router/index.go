package view_router

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/routers/middleware"
	"io/ioutil"
	"net/http"
)

func RunViewRouter(r *gin.Engine) *gin.Engine {

	//登录限制
	r.Use(middleware.HTTPAuthMiddleware())

	r.Static("/js", "./html/js")
	r.Static("/css", "./html/css")
	r.Static("/fonts", "./html/fonts")

	//index
	r.GET("/", func(c *gin.Context) {
		content, err := ioutil.ReadFile("./html/index.html")
		if err != nil {
			c.String(http.StatusOK, "index.html not found")
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Header().Add("Accept", "text/html")
		_, err = c.Writer.Write(content)
		if err != nil {
			c.String(http.StatusOK, "index.html not found")
			return
		}
		c.Writer.Flush()
	})

	return r
}
