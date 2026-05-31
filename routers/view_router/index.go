package view_router

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RunViewRouter mounts static frontend assets and index page route.
func RunViewRouter(r *gin.Engine) *gin.Engine {

	// Serve built frontend bundles.
	r.Static("/js", "./html/js")
	r.Static("/css", "./html/css")
	r.Static("/fonts", "./html/fonts")

	// Serve SPA index page.
	r.GET("/", func(c *gin.Context) {
		content, err := ioutil.ReadFile("./html/index.html")
		if err != nil {
			c.String(http.StatusOK, "index.html not found")
			return
		}
		c.Writer.WriteHeader(http.StatusOK)
		c.Writer.Header().Add("Accept", "text/html")
		// Write and flush content explicitly for compatibility.
		_, err = c.Writer.Write(content)
		if err != nil {
			c.String(http.StatusOK, "index.html not found")
			return
		}
		c.Writer.Flush()
	})

	return r
}
