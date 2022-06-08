package system

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"server/config"
)

//	Captcha 图形验证码
func (b *BaseApi) Captcha(c *gin.Context) {
	var captchaRequest config.Captcha
	_ = c.ShouldBindJSON(&captchaRequest)
	// 验证参数里面是否有验证码配置,没有读取配置文件基础设置
	//if captchaRequest.ImgWidth == 0 || captchaRequest.ImgHeight == 0 || captchaRequest.KeyLong == 0 {
	//	captchaRequest = global.Config.Captcha
	//}
	var driver base64Captcha.Driver
	switch captchaRequest.Type {
	case "audio":
		driver = base64Captcha.NewDriverAudio(captchaRequest.Length, captchaRequest.Language)
	case "string":
		driver = base64Captcha.NewDriverString(captchaRequest.Height, captchaRequest.Width, captchaRequest.NoiseCount, captchaRequest.ShowLineOptions, captchaRequest.Length, base64Captcha.TxtAlphabet, captchaRequest.BgColor, nil, []string{"wqy-microhei.ttc"})
	case "math":
		driver = base64Captcha.NewDriverMath(captchaRequest.Height, captchaRequest.Width, captchaRequest.NoiseCount, captchaRequest.ShowLineOptions, captchaRequest.BgColor, nil, nil)
	case "chinese":
		//driver = base64Captcha.NewDriverChinese()
	default:

	}
	fmt.Println(driver)
}
