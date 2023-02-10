package Email

import (
	"demo/Global"
	response "demo/Response"
	"demo/Structs"
	"demo/Utils"

	"github.com/gin-gonic/gin"
)

//用户注册
func Regist(c *gin.Context) {

	var r Structs.Regist

	if err := c.ShouldBind(&r); err != nil {
		response.FailWithDetailed(response.ParamErr, err.Error(), "提交信息非法", c)
		return
	}

	// 判断用户名是否存在
	if Utils.IsAccountExist(r.Account) {
		response.FailWithInfor(response.ParamErr, "用户名已存在", c)
		return
	}

	// 判断邮箱是否存在
	if Utils.IsEmailExist(r.Email) {
		response.FailWithInfor(response.ParamErr, "邮箱已存在", c)
		return
	}

	ToRegist(r, c)
}

func ToRegist(r Structs.Regist, c *gin.Context) {

	// 验证码验证
	if !VerifyEmail(r) {
		response.FailWithInfor(response.ServerErr, "验证码错误", c)
		return
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"Msg":  "登录注册成功",
		})
	}

	u := Structs.User{
		Account: r.Account,
		Password: r.Password,
		Email:    r.Email,
	}

	Global.DB.Create(&u)

}
