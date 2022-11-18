package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"goRedisAdmin/global/global_response"
	"goRedisAdmin/utils"
	"strconv"
	"strings"
	"time"
)

//AuthorizeJWT -> to authorize JWT Token 校验jwt token
func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := new(global_response.Response)
		const BearerSchema string = "Bearer "
		authHeader := strings.TrimSpace(ctx.GetHeader("Authorization"))
		if authHeader == "" || len(authHeader) <= 7 {
			resp.RespError("Authorization Bearer Token is empty!", ctx)
			ctx.Abort()
			return
		}
		tokenString := authHeader[len(BearerSchema):]
		if token, err := utils.BCYValidateToken(tokenString); err != nil {
			resp.RespError("Token 无效!", ctx)
			ctx.Abort()
			return
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				resp.RespError("Unauthorized!", ctx)
				ctx.Abort()
				return
			} else {
				//token解析成功
				if token.Valid {
					//自动为即将过期token生成新token
					expTime, _ := strconv.ParseInt(fmt.Sprintf("%v", claims["exp"]), 10, 64)
					now := time.Now().Unix()
					bufferTime := int64(300) //自动更新token时间间隔
					if expTime-now <= bufferTime {
						//自动更新token
						user, _ := strconv.Atoi(fmt.Sprintf("%v", claims["user"]))
						newToken, _ := utils.BCYGenerateToken(fmt.Sprintf("%v", user))
						//返回新token
						ctx.Header("new-token", newToken)
					}
					ctx.Set("user", claims["user"])
				} else {
					resp.RespError("Unauthorized!", ctx)
					ctx.Abort()
					return
				}
			}
		}
	}
}
