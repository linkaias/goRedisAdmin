package middleware

import (
	"fmt"
	"goRedisAdmin/global/global_response"
	"goRedisAdmin/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// AuthorizeJWT validates Bearer JWT from Authorization header.
//
// It returns code=NoLOGIN when token is missing/invalid.
// For valid tokens that are close to expiration, it auto-issues a new token
// and writes it into response header `new-token`.
func AuthorizeJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp := new(global_response.Response)
		const BearerSchema string = "Bearer "
		authHeader := strings.TrimSpace(ctx.GetHeader("Authorization"))
		if authHeader == "" || len(authHeader) <= 7 {
			// Reject request when Authorization header is missing.
			resp.Response(
				global_response.NoLOGIN, "Authorization Bearer Token is empty!", map[string]interface{}{}, ctx,
			)
			ctx.Abort()
			return
		}
		tokenString := authHeader[len(BearerSchema):]
		if token, err := utils.BCYValidateToken(tokenString); err != nil {
			resp.Response(
				global_response.NoLOGIN, "Token 无效!", map[string]interface{}{}, ctx,
			)
			ctx.Abort()
			return
		} else {
			if claims, ok := token.Claims.(jwt.MapClaims); !ok {
				// Parsed token did not contain expected claim structure.
				resp.Response(
					global_response.NoLOGIN, "Token 无效!Unauthorized", map[string]interface{}{}, ctx,
				)
				ctx.Abort()
				return
			} else {
				// Token parsing succeeded.
				if token.Valid {
					// Auto-refresh tokens that are near expiration.
					expTime, _ := strconv.ParseInt(fmt.Sprintf("%v", claims["exp"]), 10, 64)
					now := time.Now().Unix()
					bufferTime := int64(300) // Refresh window in seconds.
					if expTime-now <= bufferTime {
						// Reissue token and return it in response header.
						user, _ := strconv.Atoi(fmt.Sprintf("%v", claims["user"]))
						newToken, _ := utils.BCYGenerateToken(fmt.Sprintf("%v", user))
						// Return renewed token to client.
						ctx.Header("new-token", newToken)
					}
					ctx.Set("user", claims["user"])
				} else {
					// Signature/claims failed validation.
					resp.Response(
						global_response.NoLOGIN, "Token 无效!Unauthorized", map[string]interface{}{}, ctx,
					)
					ctx.Abort()
					return
				}
			}
		}
	}
}
