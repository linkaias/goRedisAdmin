package global_write_ip

// WriteListIp stores allowed client IPs as a set map (ip -> ip).
// When empty, IP filtering is treated as disabled.
var WriteListIp map[string]string

// init allocates the in-memory whitelist map.
func init() {
	WriteListIp = make(map[string]string)
}
