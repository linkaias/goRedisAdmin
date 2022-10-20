package main

import (
	"fmt"
	"goRedisAdmin/global/global_redis"
	"goRedisAdmin/global/initData"
	"goRedisAdmin/routers"
)

func init() {
	initData.Initialization()
}

func test() {
	client, err := global_redis.GetRedisClient(1)
	fmt.Println(err)
	fmt.Println(client)
}

func main() {
	//test()
	routers.RunApp()
}
