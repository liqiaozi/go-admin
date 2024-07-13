package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"lixuefei.com/go-admin/app/admin/model/dto"
	"lixuefei.com/go-admin/common/component/captcha"
	"lixuefei.com/go-admin/common/component/logger"
	"lixuefei.com/go-admin/common/errors"
	"lixuefei.com/go-admin/common/response"
	"lixuefei.com/go-admin/common/utils"
	"lixuefei.com/go-admin/global"
	"net/http"
	"time"
)

type CaptchaController struct{}

// var store = base64Captcha.DefaultMemStore
var store = captcha.NewDefaultRedisStore()

func (b CaptchaController) GenCaptcha(c *gin.Context) {
	// 判断验证码是否开启
	openCaptcha := global.App.Server.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.App.Server.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	keyLength := global.App.Server.Captcha.KeyLength                   // 验证码长度

	// 获取客户端IP
	clientIP := c.ClientIP()
	//value, err := global.App.RedisClient.Get(context.Background(), clientIP).Int()
	//if err != nil {
	//	global.App.RedisClient.Set(context.Background(), clientIP, 1, time.Second*time.Duration(openCaptchaTimeOut))
	//}
	value, ok := global.App.BlackCache.Get(clientIP)
	if !ok {
		global.App.BlackCache.Set(clientIP, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}
	var openCaptchaFlag bool
	if openCaptcha == 0 || openCaptcha < utils.InterfaceToInt(value) {
		openCaptchaFlag = true
	}

	// 字符,公式,验证码配置,生成默认数字的driver
	imgHeight := global.App.Server.Captcha.ImgHeight
	imgWidth := global.App.Server.Captcha.ImgWidth
	driver := base64Captcha.NewDriverDigit(imgHeight, imgWidth, keyLength, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, _, err := cp.Generate()
	if err != nil {
		logger.Log.Errorf("gen captcha err: %s", err.Error())
		c.JSON(http.StatusOK, response.FailByError(errors.CaptchaError))
		return
	}

	// 返回验证码
	captchaInfo := dto.CaptchaDTO{
		CaptchaId:     id,
		PicPath:       b64s,
		CaptchaLength: keyLength,
		OpenCaptcha:   openCaptchaFlag,
	}
	c.JSON(http.StatusOK, response.OkWithData(captchaInfo))
}
