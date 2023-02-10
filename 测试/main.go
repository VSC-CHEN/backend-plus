package main

import (
	"demo/MiddleWare/Cross"
	"demo/Connect/MySQL"
	"demo/Viper"
	"demo/Login"

	"demo/Email"
	"demo/Connect/Redis"

	"github.com/gin-gonic/gin"
)

func init() {
	Viper.Config()
	MySQL.Mysql()
	Redis.Redis()
}

func main() {
	r := gin.Default()
	WOToken := r.Group("/user")
	WOToken.Use(Cross.Cors())
	{
		//登录
		WOToken.POST("/login", Login.Login)
		//注册
		WOToken.POST("/regist", Email.Regist)
		//邮件发送
		WOToken.POST("/email", Email.SendEmailCode)
	}

	r.Run(":8000")
}
