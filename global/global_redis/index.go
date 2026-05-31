package global_redis

import (
	"errors"

	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
)

// rdConfigStruct stores Redis connection options loaded from config.ini.
type rdConfigStruct struct {
	Name      string
	Port      string
	Host      string
	Pwd       string
	Timeout   int
	DoTimeout int
}

// cfg is the singleton runtime Redis config.
var cfg = new(rdConfigStruct)

// SetRDConfig copies Redis config values from INI section into memory.
func SetRDConfig(config *ini.Section) {
	cfg.Name = config.Key("name").String()
	cfg.Port = config.Key("port").String()
	cfg.Host = config.Key("host").String()
	cfg.Pwd = config.Key("pwd").String()
	cfg.Timeout, _ = config.Key("timeout").Int()
	cfg.DoTimeout, _ = config.Key("do_timeout").Int()
}

// GetRedisClient creates and verifies a Redis client for a specific DB index.
//
// Constraints:
//   - Redis DB must be in [0, 15].
//
// Behavior:
//   - Builds a new client each call.
//   - Executes Ping to ensure connectivity before returning.
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
