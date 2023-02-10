package Email

import (
	response "demo/Response"
	"demo/Structs"

	"github.com/gin-gonic/gin"
)

//注册账户邮箱的验证码
func SendEmailCode(c *gin.Context) {

	var e Structs.Emails

	// 读取数据
	if err := c.ShouldBind(&e); err != nil {
		response.FailWithDetailed(response.ParamErr, err.Error(), "提交信息非法", c)
		return
	}

	// 发送邮箱验证码
	SendCode(e, c)
}
