package middlewares

import (
	"G02-Go-API/pkg/jwt"
	"G02-Go-API/pkg/response"
	"github.com/gin-gonic/gin"
)

// GuestJWT 强制使用游客身份访问
func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		if len(c.GetHeader("Authorization")) > 0 {

			// 解析 token 成功，说明登录成功了
			_, err := jwt.NewJWT().ParserToken(c)
			if err == nil {
				response.Unauthorized(c, "请使用游客身份登录")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
