// Package routes 注册路由
package routes

import (
	"G02-Go-API/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放在这里
	v1 := r.Group("/v1")
	{
		//// 注册一个路由
		//v1.GET("/", func(c *gin.Context) {
		//	// 以 JSON 格式响应
		//	c.JSON(http.StatusOK, gin.H{
		//		"Hello": "World!",
		//	})
		//})

		// 测试一个 v1 的路由器，我们所有的 v1 版本的路由都将存放在这里
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已经注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			//判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExits)
		}
	}
}
