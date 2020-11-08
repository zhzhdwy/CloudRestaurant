package controller

import (
	"CloudRestaurant/model"
	"CloudRestaurant/param"
	"CloudRestaurant/service"
	"CloudRestaurant/tool"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

type MemberController struct {
}

func (mc *MemberController) Router(engine *gin.Engine) {
	// http://localhost:8090/api/sendcode?phone=18611441751
	engine.GET("/api/sendcode", mc.sendSmsCode)
	// http://localhost:8090/api/login_sms
	engine.POST("/api/login_sms", mc.smsLogin)
	// http://localhost:8090/api/captcha
	engine.GET("/api/captcha", mc.captcha)
	// 用户头像上传
	engine.POST("/api/vertifycha", mc.vertifyCaptcha)
	//用户名密码登陆
	engine.POST("/api/login_pwd", mc.nameLogin)
	//头像上传
	engine.POST("/api/upload/avator", mc.uploadAvator)
	//用户信息查询
	engine.GET("/api/userinfo", mc.userInfo)

}

func (mc *MemberController) userInfo(context *gin.Context) {
	cookie, err := tool.CookieAuth(context)
	if err != nil {
		context.Abort()
		tool.Failed(context, "还未登陆，请先登陆")
		return
	}
	memberService := service.MemberService{}
	member := memberService.GetUserInfo(cookie.Value)
	if member != nil {
		//这地方可以把password去了
		tool.Success(context, &member)
		return
	}
	tool.Failed(context, "获取用户信息失败")
}

func (mc *MemberController) uploadAvator(context *gin.Context) {
	userId := context.PostForm("user_id")
	fmt.Println(userId)
	file, err := context.FormFile("avatar")
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}

	//判断用户是否登陆
	sess := tool.GetSess(context, "user_"+userId)
	if sess != nil {
		tool.Failed(context, "用户没有登陆")
		return
	}
	var member model.Member
	json.Unmarshal(sess.([]byte), &member)

	// 保存file
	fileName := "./uploadfile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = context.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.Failed(context, "头像更新失败")
		return
	}

	//将文件保存到fastdfs中
	fileId := tool.UploadFile(fileName)
	if fileId != "" {
		//删除本地文件
		os.Remove(fileName)

		//保存本地文件目录到数据库的头像字段中
		memberService := service.MemberService{}
		path := memberService.UploadAvatar(member.Id, fileId)
		if path != "" {
			tool.Success(context, tool.FileServerAddr()+"/"+path)
			return
		}

	}

	tool.Failed(context, "上传失败")
}

func (mc *MemberController) nameLogin(context *gin.Context) {
	//获取参数
	var loginParam param.LoginParam
	err := tool.Decode(context.Request.Body, &loginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}

	//验证验证码
	captcha := tool.VertifyCaptcha(loginParam.Id, loginParam.Value)
	if !captcha {
		tool.Failed(context, "验证码不正确，请重新输入")
		return
	}
	//登陆
	ms := service.MemberService{}
	member := ms.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0 {
		// 保持会话
		sess, _ := json.Marshal(member)
		err := tool.Setsess(context, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Failed(context, "登陆失败")
			return
		}
		tool.Success(context, &member)
		return
	}
	tool.Failed(context, "登陆失败")

}

// 校验验证码
func (mc *MemberController) vertifyCaptcha(context *gin.Context) {
	var captcha tool.CaptchaResult
	err := tool.Decode(context.Request.Body, &captcha)
	if err != nil {
		tool.Failed(context, "参数解析失败")
	}
	result := tool.VertifyCaptcha(captcha.Id, captcha.VertifyValue)
	if result {
		fmt.Println("验证通过")
	} else {
		fmt.Println("验证失败")
	}
}

//生产验证码
func (mc *MemberController) captcha(context *gin.Context) {
	tool.GenerateCaptcha(context)
}

// http://localhost:8090/api/sendcode?phone=18516261623
func (mc *MemberController) sendSmsCode(context *gin.Context) {
	//发送验证码
	phone, exist := context.GetQuery("phone")
	if !exist {
		tool.Failed(context, "参数解析失败")
		return
	}

	ms := service.MemberService{}
	isSend := ms.Sendcode(phone)
	if isSend {
		tool.Success(context, "发送成功")
		return
	}
	tool.Failed(context, "发送失败")
}

//手机号+短信登陆方式
func (mc *MemberController) smsLogin(context *gin.Context) {
	var smsLoginParam param.SmsLoginParam

	fmt.Println(context.PostForm("phone"))
	err := tool.Decode(context.Request.Body, &smsLoginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
	}
	defer context.Request.Body.Close()
	//完成手机+验证码登陆
	us := service.MemberService{}
	member := us.SmsLogin(smsLoginParam)
	if member != nil {
		sess, _ := json.Marshal(member)
		err := tool.Setsess(context, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Failed(context, "登陆失败")
			return
		}
		// 设置cookie，和session保存任选一
		context.SetCookie("cookie_user", strconv.Itoa(int(member.Id)), 10*60,
			"/", "localhost", true, true)
		tool.Success(context, member)
		return
	}
	tool.Failed(context, "登陆失败")

}
