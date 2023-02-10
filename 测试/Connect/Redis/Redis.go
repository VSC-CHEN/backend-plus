package Redis

import (
	"demo/Global"
	"log"

	"github.com/go-redis/redis"
)

func Redis() {
	// 获取 redis 配置
	r := Global.Config.Redis
	// 建立连接
	client := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})
	// 测试连接
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal("连接redis数据库失败: ", err)
	} else {
		Global.REDIS_DB = client
		log.Print("成功创建redis数据库: ", pong)
	}
}
