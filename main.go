package main

import (
	"goRedisAdmin/global/initData"
	"goRedisAdmin/routers"
	"goRedisAdmin/utils/log_utils"
)

// init loads configuration and initializes global runtime data
// (Redis config, IP whitelist, etc.) before main starts.
func init() {
	initData.Initialization()
}

// main starts background log workers and then boots the HTTP server.
func main() {
	// Start async log consumers.
	log_utils.RunLog()
	// Start Gin application and block.
	routers.RunApp()
}
