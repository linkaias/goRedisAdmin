package db_data_controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goRedisAdmin/controller"
	"goRedisAdmin/global/global_redis"
)

type dbDataCont struct {
	controller.BaseController
}

func NewDbDataController() DbDataController {
	cont := new(dbDataCont)
	cont.BaseInit()
	return cont
}

type DbDataController interface {
	DbList(ctx *gin.Context)
	GetKeys(ctx *gin.Context)
}

func (c dbDataCont) DbList(ctx *gin.Context) {
	data := make([]map[string]interface{}, 0)
	for i := 0; i < 16; i++ {
		temp := make(map[string]interface{})
		num, _ := getDbKeyLen(i)
		temp["db_num"] = i
		temp["keys_len"] = num
		temp["show_name"] = fmt.Sprintf("Db%v", i)
		data = append(data, temp)
	}
	c.Resp.RespSuccessWithData(data, ctx)
}
func getDbKeyLen(i int) (uint, error) {
	rd, err := global_redis.GetRedisClient(i)
	if err != nil {
		return 0, err
	}
	defer rd.Close()
	num := rd.DbSize().Val()
	return uint(num), nil
}

// GetKeys 获取数据库的keys
func (c dbDataCont) GetKeys(ctx *gin.Context) {
	//default value is 0
	dbNum, _ := c.ParamToInt(ctx, "db_num", "get")
	rd, err := global_redis.GetRedisClient(dbNum)
	if err != nil {
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer rd.Close()
	filter := "*"
	info, _ := rd.Keys(filter).Result()
	data := make([]map[string]interface{}, 0)
	for i, key := range info {
		temp := make(map[string]interface{})
		temp["id"] = i + 1
		keyType, err := rd.Type(key).Result()
		if err != nil {
			temp["msg"] = "Error:" + err.Error()
		} else {
			temp["msg"] = "Success!"
		}
		temp["type"] = keyType
		temp["key"] = key
		expire, _ := rd.TTL(key).Result()
		expireS := expire.Seconds()
		if expireS == -1 {
			temp["expire_at"] = "长期"
		} else {
			temp["expire_at"] = expire.String()
		}
		data = append(data, temp)
	}
	c.Resp.RespSuccessWithData(data, ctx)
}
