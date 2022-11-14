package middleware

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/global/global_write_ip"
	"goRedisAdmin/global/initData"
	"net/http"
)

//HTTPAuthMiddleware login middleware
func HTTPAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cfg := initData.IniRead.Section("admin")
		user := cfg.Key("username").String()
		passwd := cfg.Key("password").String()
		reqUser, reqPasswd, hasAuth := ctx.Request.BasicAuth()
		if (user == "" && passwd == "") ||
			(hasAuth && reqUser == user && reqPasswd == passwd) {
			ctx.Next()
		} else {
			ctx.Writer.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(ctx.Writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			ctx.Abort()
		}
	}
}

func IpCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(global_write_ip.WriteListIp) > 0 {
			ip := c.ClientIP()
			_, ok := global_write_ip.WriteListIp[ip]
			if !ok {
				c.JSON(http.StatusBadGateway, "illegal ip!")
				c.Abort()
				return
			}
		}
		// 请求前
		c.Next()
		// 请求后
	}
}
