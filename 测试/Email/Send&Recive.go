package Email

import (
	"demo/Global"
	response "demo/Response"
	"demo/Structs"
	"demo/Utils"
	"fmt"
	"log"
	"net/smtp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)

//发送验证码
func SendCode(e Structs.Emails, c *gin.Context) {

	key := e.Email + ".send"

	// 是否超过频率
	err := Global.REDIS_DB.Get(key).Err()
	if  err == nil {
		response.FailWithInfor(response.ParamErr, "验证码发送过于频繁", c)
		return
	}

	// 生成验证码
	code := Utils.RandomString(6)
	// 验证码过期时间
	expires := Global.Config.Email.Expires
	// 构造邮件的内容
	mail := &email.Email{
		From:    fmt.Sprintf("%v <%v>", Global.Config.Email.Name, Global.Config.Email.Addr),
		To:      []string{e.Email},
		Subject: "邮箱验证",
		Text:    []byte("【考核测试】验证码：" + code + "，有效期 " + strconv.Itoa(expires) + " 秒"),
	}

	// 发送邮件
	err1 := mail.Send(Global.Config.Email.Smtp, smtp.PlainAuth("", Global.Config.Email.Addr, Global.Config.Email.Password, Global.Config.Email.Host))
	if  err1 != nil {
		log.Print("邮件发送失败: ", err1)

		response.FailWithInfor(response.ParamErr, "发送失败, 请确认邮箱是否可用", c)

	} else {
		// 重新发送限制
		Global.REDIS_DB.Set(key, 1, time.Second*60)
		// 存储验证码
		Global.REDIS_DB.Set(e.Email, code, 60*time.Second)

		response.SucceedWithInfor("发送成功", c)
	}
}

//验证验证码
func VerifyEmail(r Structs.Regist) bool {

	//管理员测试码
	if gin.Mode() == gin.DebugMode {
		if r.VerifyCode == "111111" {
			return true
		}
	}

	// 获取 Redis 中的验证码
	redisCode, err := Global.REDIS_DB.Get(r.Email).Result()
	if err != nil {
		return false
	}
	// 比较验证码
	return r.VerifyCode == redisCode
}
