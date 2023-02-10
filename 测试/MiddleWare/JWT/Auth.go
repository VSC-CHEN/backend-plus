package JWT

import (
	response "demo/Response"
	"demo/Utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取 Token
		tokenString := c.GetHeader("X-Token")
		if tokenString == "" {
			response.FailWithInfor(response.NoLogin, "未登录", c)
			c.Abort()
			return
		}

		// 解析 Token
		token, claims, err := Utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			response.FailWithInfor(response.TokenExpired, "登录已过期", c)
			c.Abort()
			return
		}

		// 写入上下文
		c.Set("UID", claims.UID)
		c.Set("Account", claims.Account)
		c.Set("Password", claims.Password)
		c.Next()
	}
}
