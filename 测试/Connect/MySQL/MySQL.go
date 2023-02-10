package MySQL

import (
	"demo/Global"
	"demo/Structs"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Mysql() {
	// 获取 mysql 配置
	mCfg := Global.Config.Mysql.Dsn()
	// 连接
	if db, err := gorm.Open(mysql.Open(mCfg)); err != nil {
		log.Fatal("Connect mysql failed: ", err)
	} else {
		// 自动建表
		db.AutoMigrate(&Structs.User{})
		Global.DB = db
		log.Print("Connect mysql success: ", mCfg)
	}
}
