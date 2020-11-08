package tool

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type CaptchaResult struct {
	Id           string `json:"id"`
	Base64blob   string `json:"base_64_blob"`
	VertifyValue string `json:"vertify_value"`
}

//生产验证码图像
func GenerateCaptcha(ctx *gin.Context) {
	//base64Captcha.ConfigCharacter 生产验证码
	configCharacter := base64Captcha.ConfigCharacter{
		Height:                 30,
		Width:                  60,
		Mode:                   3,
		IsUseSimpleFont:        true,
		ComplexOfNoiseText:     0,
		ComplexOfNoiseDot:      0,
		IsShowHollowLine:       false,
		IsShowNoiseDot:         false,
		IsShowNoiseText:        false,
		IsShowSlimeLine:        false,
		IsShowSineLine:         false,
		ChineseCharacterSource: "",
		SequencedCharacters:    nil,
		UseCJKFonts:            false,
		CaptchaLen:             4,
		BgHashColor:            "",
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 254,
		},
	}

	// id和实例
	captchaId, captchaInterfaceInstace := base64Captcha.GenerateCaptcha("", configCharacter)
	base64blob := base64Captcha.CaptchaWriteToBase64Encoding(captchaInterfaceInstace)

	CaptchaResult := CaptchaResult{
		Id:         captchaId,
		Base64blob: base64blob,
	}

	// 通过success返回给前端即可
	Success(ctx, map[string]interface{}{
		"captcha_result": CaptchaResult,
	})
}

func VertifyCaptcha(id string, value string) bool {
	captcha := base64Captcha.VerifyCaptcha(id, value)
	return captcha
}
