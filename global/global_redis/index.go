package global_redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
)

type rdConfigStruct struct {
	Name      string
	Port      string
	Host      string
	Pwd       string
	Timeout   int
	DoTimeout int
}

var cfg = new(rdConfigStruct)

func SetRDConfig(config *ini.Section) {
	cfg.Name = config.Key("name").String()
	cfg.Port = config.Key("port").String()
	cfg.Host = config.Key("host").String()
	cfg.Pwd = config.Key("pwd").String()
	cfg.Timeout, _ = config.Key("timeout").Int()
	cfg.DoTimeout, _ = config.Key("do_timeout").Int()
	fmt.Println(cfg)
}

func GetRedisClient(db int) (*redis.Client, error) {
	if db < 0 || db > 15 {
		return nil, errors.New("db must be between 0 and 15")
	}
	rdb := redis.NewClient(&redis.Options{Addr: cfg.Host + ":" + cfg.Port, Password: cfg.Pwd, DB: db})
	_, err := rdb.Ping().Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
