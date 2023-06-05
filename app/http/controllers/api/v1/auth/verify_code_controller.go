package auth

import (
	v1 "G02-Go-API/app/http/controllers/api/v1"
	"G02-Go-API/pkg/captcha"
	"G02-Go-API/pkg/logger"
	"G02-Go-API/pkg/response"
	"github.com/gin-gonic/gin"
)

// VerifyCodeController 用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}

// ShowCaptcha 显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	// 记录错误日志， 因为验证码是用户的入口， 出错时应该记 error 级别的日记
	logger.LogIf(err)
	// 返回给用户
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
