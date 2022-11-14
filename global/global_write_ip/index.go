package global_write_ip

// WriteListIp ip 白名单
var WriteListIp map[string]string

func init() {
	WriteListIp = make(map[string]string)
}
