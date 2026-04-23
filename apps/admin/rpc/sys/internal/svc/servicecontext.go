package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"gorm.io/gorm"
	"sys/core"
	"sys/internal/config"
	"sys/internal/models"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	Redis    *redis.Redis
	RedisKey string // redis的模块统一前缀
}

func NewServiceContext(c config.Config) *ServiceContext {
	redisKey := c.RpcServerConf.Redis.Key
	rds := redis.MustNewRedis(c.RpcServerConf.Redis.RedisConf)
	db, err := core.InitDB(&c)
	if err != nil {
		panic(err)
	}
	models.InitModels(db)
	return &ServiceContext{
		Config:   c,
		DB:       db,
		Redis:    rds,
		RedisKey: redisKey,
	}
}
