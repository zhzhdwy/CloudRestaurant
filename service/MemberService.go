package service

import (
	"CloudRestaurant/dao"
	"CloudRestaurant/model"
	"CloudRestaurant/param"
	"CloudRestaurant/tool"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"math/rand"
	"strconv"
	"time"
)

type MemberService struct {
}

func (ms *MemberService) GetUserInfo(userId string) *model.Member {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil
	}
	memberDao := dao.MemberDao{tool.DbEngine}
	return memberDao.QueryMemberById(int64(id))
}

func (ms *MemberService) UploadAvatar(userId int64, fileName string) string {
	memberDao := dao.MemberDao{tool.DbEngine}
	result := memberDao.UpdateMemberAvatar(userId, fileName)
	if result == 0 {
		return ""
	}
	return fileName
}

//用户登陆
func (ms *MemberService) Login(name string, password string) *model.Member {

	//用户存在直接返回
	md := dao.MemberDao{tool.DbEngine}
	member := md.Query(name, password)
	if member.Id != 0 {
		return member
	}

	user := model.Member{}
	user.Username = name
	user.Password = tool.EncoderSha256(password)
	user.RegisterTime = time.Now().Unix()

	result := md.InsertMember(user)
	user.Id = result

	return &user

}

//用户手机号验证码登陆
func (ms *MemberService) SmsLogin(loginparam param.SmsLoginParam) *model.Member {
	md := dao.MemberDao{tool.DbEngine}
	sms := md.ValidateSmsCode(loginparam.Phone, loginparam.Code)
	if sms.Id == 0 {
		return nil
	}

	member := md.QueryByPhone(loginparam.Phone)
	if member.Id != 0 {
		return member
	}

	user := model.Member{}
	user.Username = loginparam.Phone
	user.Mobile = loginparam.Phone
	user.RegisterTime = time.Now().Unix()

	user.Id = md.InsertMember(user)

	return &user
}

func (ms *MemberService) Sendcode(phone string) bool {
	//产生验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	//调用阿里云sdk
	//config := tool.GetConfig().Sms
	//client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
	//if err != nil {
	//	logger.Error(err.Error())
	//	return false
	//}
	//request := dysmsapi.CreateSendSmsRequest()
	//request.Scheme = "https"
	//request.SignName = config.SignName
	//request.TemplateCode = config.TemplateCode
	//request.PhoneNumbers = phone
	//par, err := json.Marshal(map[string]interface{}{
	//	"code": code,
	//})
	//request.TemplateParam = string(par)
	//response, err := client.SendSms(request)
	//if err != nil {
	//	logger.Error(err.Error())
	//	return false
	//}
	// response手动定义，测试数据库
	response := new(dysmsapi.SendSmsResponse)
	response.Message = "OK"
	response.RequestId = "111-111-111"
	response.BizId = "123123143512435"
	response.Code = "OK"
	fmt.Println(response)
	// 短信发送成功
	if response.Code == "OK" {
		//验证码保存到数据库中
		smsCode := model.SmsCode{
			Phone:      phone,
			Code:       code,
			BizId:      response.BizId,
			CreateTime: time.Now().Unix(),
		}
		memberDao := dao.MemberDao{tool.DbEngine}
		result := memberDao.InsertCode(smsCode)
		return result > 0
	}

	//接受判断发送状态
	return false
}
