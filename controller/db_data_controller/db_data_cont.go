package db_data_controller

import (
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
	data := make([]uint, 16)
	for i := 0; i < 16; i++ {
		num, _ := getDbKeyLen(i)
		data[i] = num
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

	c.Resp.RespSuccessWithData(info, ctx)
}
