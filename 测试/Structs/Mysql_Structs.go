package Structs

import (
	"gorm.io/gorm"
)

type ToUser struct {
	gorm.Model
	IDCard   uint
	Name     string
	Portrait string
}

type Mysql struct {
	Host     string `mapstructure:"host"`     // 服务器地址
	Port     string `mapstructure:"port"`     // 端口
	Username string `mapstructure:"username"` // 数据库用户名
	Password string `mapstructure:"password"` // 数据库密码
	Database string `mapstructure:"database"` // 数据库名
	Config   string `mapstructure:"config"`   // 高级配置
}

type Redis struct {
	Addr     string `mapstructure:"addr"`     // 服务器地址
	Password string `mapstructure:"password"` // 数据库密码
	DB       int    `mapstructure:"db"`       // 数据库
	Method   string `mapstructure:"method"`
}

type Common struct {
	Origin string `mapstructure:"origin"`
	Port   string `mapstructure:"port"`
}

type Config struct {
	Common Common `mapstructure:"common"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Redis  Redis  `mapstructure:"redis"`
	Email  Email  `mapstructure:"email"`
	JWT    JWT    `mapstructure:"jwt"`
}

//表内部包含信息
type User struct {
	gorm.Model
	UID      uint
	Account  string
	Password string 
	Email    string
}
//返回校验信息
func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Database + "?" + m.Config
}
