package config

import (
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		Host     string
		Port     int
		User     string
		Password string
		Source   string
		DB       string
	}

	JWT struct {
		AccessSecret string
		AccessExpire int64
	}
}

func (d *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.Mysql.User, d.Mysql.Password, d.Mysql.Host, d.Mysql.Port, d.Mysql.DB)
}
func (d *Config) Empty() bool {
	return d.Mysql.User == "" && d.Mysql.Password == "" && d.Mysql.Host == "" && d.Mysql.Port == 0 && d.Mysql.DB == ""
}
