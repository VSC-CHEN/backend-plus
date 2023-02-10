package Login

import (
	"demo/Global"
	response "demo/Response"
	"demo/Structs"
	"demo/Utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

//用户登录
func Login(c *gin.Context) {
	var user Structs.User

	Account := c.PostForm("Account")
	Password := c.PostForm("Password")

	result := Global.DB.Where("account = ?", Account).First(&user)

	TokenString, err := Utils.ReleaseToken(user)

	if result.Error != nil || user.Password != Password {
		response.FailWithInfor(response.NoAuth, "密码错误", c)
	} else if user.Account == "" {
		response.FailWithInfor(response.ServerErr, "账户不存在", c)
	} else if err != nil {
		response.FailWithInfor(response.ServerErr, "Token生成失败请联系管理员", c)
	} else {
		response.Result(response.Success, fmt.Sprintf("Token:%s", TokenString), fmt.Sprintf("登录成功:%s", Account), c)
	}
}