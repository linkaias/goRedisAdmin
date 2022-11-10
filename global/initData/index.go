package initData

import (
	"goRedisAdmin/global/global_redis"
	"gopkg.in/ini.v1"
	"log"
)

//IniRead ini reader
var IniRead *ini.File

func init() {
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		log.Panicln(err)
	}
	IniRead = cfg
}

func Initialization() {
	global_redis.SetRDConfig(IniRead.Section("redis"))
}
