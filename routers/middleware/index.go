package middleware

import (
	"goRedisAdmin/global/global_write_ip"
	"goRedisAdmin/global/initData"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPAuthMiddleware provides optional HTTP Basic Auth protection.
//
// It compares incoming credentials with [admin] username/password from config.
// If admin credentials are empty, the middleware allows all requests.
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

// IpCheckMiddleware blocks requests whose client IP is not in whitelist.
//
// Behavior:
//   - If whitelist is empty: allow all IPs.
//   - If whitelist has entries: allow only exact matching client IPs.
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
		// Continue to next handler.
		c.Next()
		// Post-handler hook reserved for future use.
	}
}
