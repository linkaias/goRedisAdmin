package exoprt_utils

import (
	"goRedisAdmin/global/global_redis"
	"testing"
)

func TestExportUtils_ExportFile(t *testing.T) {
	cont := &ExportUtils{}
	rd, _ := global_redis.GetRedisClient(0)

	cont.LoadKeysData(rd, []string{"fffff", "tttt"})
	cont.ExportFile()
}
