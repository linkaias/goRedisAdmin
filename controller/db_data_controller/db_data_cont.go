package db_data_controller

import (
	"errors"
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
	GetVal(ctx *gin.Context)
	DelKey(ctx *gin.Context)
	AddVal(ctx *gin.Context)
	Flush(ctx *gin.Context)
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
			temp["expire_at"] = "长期有效"
		} else {
			temp["expire_at"] = expire.String()
		}
		data = append(data, temp)
	}
	c.Resp.RespSuccessWithData(data, ctx)
}

// GetVal 获取详细的value
func (c dbDataCont) GetVal(ctx *gin.Context) {
	s := new(DbDataHelpModel)
	err := ctx.ShouldBind(s)
	if err != nil {
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	if s.VType == "" {
		c.Resp.RespError("type is required", ctx)
		return
	}
	cont, err := NewDbDataHelpController(s)
	if err != nil {
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer cont.CloseClient()
	res, err := handleGetVal(s.VType, cont)
	if err != nil {
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	c.Resp.RespSuccessWithData(res, ctx)

}

func (c dbDataCont) Flush(ctx *gin.Context) {
	dbNum, _ := c.ParamToInt(ctx, "db_num", "get")
	rd, err := global_redis.GetRedisClient(dbNum)
	if err != nil {
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer rd.Close()
	flushType := ctx.Query("type")
	if flushType == "db" {
		_, err = rd.FlushDB().Result()
		if err != nil {
			c.Resp.RespError(err.Error(), ctx)
			return
		}
	} else if flushType == "all" {
		_, err = rd.FlushAll().Result()
		if err != nil {
			c.Resp.RespError(err.Error(), ctx)
			return
		}
	} else {
		c.Resp.RespError("type is required", ctx)
		return
	}
	c.Resp.RespSuccess(ctx)
}

// DelKey 删除key
func (c dbDataCont) DelKey(ctx *gin.Context) {
	//default db is 0
	dbNum, _ := c.ParamToInt(ctx, "db_num", "get")
	rd, err := global_redis.GetRedisClient(dbNum)
	if err != nil {
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer rd.Close()
	key := ctx.Query("key")
	if key == "" {
		c.Resp.RespError("key is required", ctx)
		return
	}
	_, err = rd.Del(key).Result()
	if err != nil {
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	c.Resp.RespSuccess(ctx)
}

// AddVal 新增键值
func (c dbDataCont) AddVal(ctx *gin.Context) {
	s := new(DbDataHelpModel)
	err := ctx.ShouldBind(s)
	if err != nil {
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	if s.VType == "" {
		c.Resp.RespError("type is required", ctx)
		return
	}
	cont, err := NewDbDataHelpController(s)
	if err != nil {
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer cont.CloseClient()
	err = handleAddVal(s.VType, cont)
	if err != nil {
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	c.Resp.RespSuccess(ctx)
}

func handleAddVal(valType string, cont *DbDataHelpCont) error {
	switch valType {
	case "string":
		return cont.AddString()
	case "list":
		return cont.AddList()
	case "set":
		return cont.AddSet()
	case "zset":
		return cont.AddZSet()
	case "hash":
		return cont.AddHash()
	}
	return errors.New("type not supported ! ")
}

func handleGetVal(valType string, cont *DbDataHelpCont) (interface{}, error) {
	switch valType {
	case "string":
		return cont.GetString()
	case "list":
	case "set":
	case "zset":
	case "hash":
	}
	return nil, errors.New("type not supported ! ")
}
