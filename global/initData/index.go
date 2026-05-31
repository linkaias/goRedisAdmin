package initData

import (
	"goRedisAdmin/global/global_redis"
	"goRedisAdmin/global/global_write_ip"
	"log"
	"strings"

	"gopkg.in/ini.v1"
)

// IniRead is the global handle of config.ini loaded at package init time.
var IniRead *ini.File

// init loads config.ini from project root; panic on failure to prevent
// starting with invalid runtime configuration.
func init() {
	cfg, err := ini.Load("./config.ini")
	if err != nil {
		log.Panicln(err)
	}
	IniRead = cfg
}

// Initialization wires global runtime config into corresponding packages.
func Initialization() {
	global_redis.SetRDConfig(IniRead.Section("redis"))
	// Initialize in-memory IP whitelist set.
	initWriteListIp(IniRead.Section("whitelist_ip"))
}

// initWriteListIp parses comma-separated allow_ip and fills whitelist map.
func initWriteListIp(config *ini.Section) {
	ips := config.Key("allow_ip").String()
	if ips != "" {
		ipStr := strings.Split(strings.TrimSpace(ips), ",")
		for _, ip := range ipStr {
			global_write_ip.WriteListIp[ip] = ip
		}
	}
}
