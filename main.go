package main

import (
	"goRedisAdmin/global/initData"
	"goRedisAdmin/routers"
	"goRedisAdmin/utils/log_utils"
)

func init() {
	initData.Initialization()
}

func main() {
	//记录日志
	log_utils.RunLog()
	//run app
	routers.RunApp()
}
