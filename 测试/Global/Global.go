package Global

import (
	"demo/Structs"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	Config   Structs.Config
	DB       *gorm.DB
	REDIS_DB *redis.Client
)
