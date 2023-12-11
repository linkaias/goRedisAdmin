package db_data_controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"goRedisAdmin/controller"
	"goRedisAdmin/global/global_redis"
	"goRedisAdmin/utils/log_utils"
	"strings"
	"time"
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
	GetValByKey(ctx *gin.Context)
	// ExpireKey 修改key的过期时间
	ExpireKey(ctx *gin.Context)
}

func (c dbDataCont) DbList(ctx *gin.Context) {
	data := make([]map[string]interface{}, 0)
	for i := 0; i < 16; i++ {
		temp := make(map[string]interface{})
		num, err := getDbKeyLen(i)
		if err != nil {
			log_utils.WriteLog("err", err, nil)
		}
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
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer rd.Close()
	filter := "*"
	query := ctx.Query("filter")
	if query != "" {
		filter = query
	}

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
		temp["len"] = getLenByKey(rd, key, keyType)

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
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	if s.VType == "" {
		c.Resp.RespError("type is required", ctx)
		return
	}
	cont, err := NewDbDataHelpController(s)
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer cont.CloseClient()
	res, err := handleGetVal(s.VType, cont)
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	c.Resp.RespSuccessWithData(res, ctx)

}

func (c dbDataCont) Flush(ctx *gin.Context) {
	dbNum, _ := c.ParamToInt(ctx, "db_num", "get")
	rd, err := global_redis.GetRedisClient(dbNum)
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer rd.Close()
	flushType := ctx.Query("type")
	if flushType == "db" {
		_, err = rd.FlushDB().Result()
		if err != nil {
			log_utils.WriteLog("err", err, nil)
			c.Resp.RespError(err.Error(), ctx)
			return
		}
	} else if flushType == "all" {
		_, err = rd.FlushAll().Result()
		if err != nil {
			log_utils.WriteLog("err", err, nil)
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
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer rd.Close()
	key := ctx.Query("key")
	if key == "" {
		c.Resp.RespError("key is required", ctx)
		return
	}
	waitDelKey := make([]string, 0)
	if strings.Contains(key, ",") {
		waitDelKey = strings.Split(key, ",")
	} else {
		waitDelKey = append(waitDelKey, key)
	}
	for _, k := range waitDelKey {
		_, err = rd.Del(k).Result()
	}
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	c.Resp.RespSuccess(ctx)
}

// ExpireKey 修改key的过期时间
func (c dbDataCont) ExpireKey(ctx *gin.Context) {
	//default db is 0
	dbNum, _ := c.ParamToInt(ctx, "db_num", "get")
	expireInt, _ := c.ParamToInt(ctx, "expire", "get")
	rd, err := global_redis.GetRedisClient(dbNum)
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer rd.Close()
	key := ctx.Query("key")
	if key == "" {
		c.Resp.RespError("key is required", ctx)
		return
	}
	if expireInt == 0 { // 设置无过期时间
		rd.Persist(key)
	} else { // 设置过期时间
		_, err = rd.Expire(key, time.Duration(expireInt)*time.Second).Result()
	}
	if err != nil {
		log_utils.WriteLog("err", err, nil)
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
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	if s.VType == "" {
		c.Resp.RespError("type is required", ctx)
		return
	}
	cont, err := NewDbDataHelpController(s)
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer cont.CloseClient()
	err = handleAddVal(s.VType, cont)
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	c.Resp.RespSuccess(ctx)
}

// GetValByKey 获取数据通过key
func (c dbDataCont) GetValByKey(ctx *gin.Context) {
	//default value is 0
	dbNum, _ := c.ParamToInt(ctx, "db_num", "get")
	rd, err := global_redis.GetRedisClient(dbNum)
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer rd.Close()
	type GetValByKeyStruct struct {
		Key   string `json:"key"`
		DType string `json:"type"`
	}
	s := new(GetValByKeyStruct)
	err = ctx.ShouldBind(s)
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	if s.Key == "" || s.DType == "" {
		c.Resp.RespError("type or key  is required", ctx)
		return
	}
	if s.DType == "string" {
		result, err := rd.Get(s.Key).Result()
		if err != nil {
			log_utils.WriteLog("err", err, nil)
			c.Resp.RespError(err.Error(), ctx)
			return
		}
		c.Resp.RespSuccessWithData(result, ctx)
		return
	}
	if s.DType == "hash" {
		//default value is 0
		cursorNum, _ := c.ParamToInt(ctx, "cursor", "get")
		filter := ctx.Query("filter")
		strings, cursor, err := rd.HScan(s.Key, uint64(cursorNum), filter, 100).Result()
		if err != nil {
			log_utils.WriteLog("err", err, nil)
			c.Resp.RespError(err.Error(), ctx)
			return
		}
		data, err := descHashPageData(strings)
		if err != nil {
			log_utils.WriteLog("err", err, nil)
			c.Resp.RespError(err.Error(), ctx)
			return
		}

		count, err := rd.HLen(s.Key).Result()
		if err != nil {
			log_utils.WriteLog("err", err, nil)
			c.Resp.RespError(err.Error(), ctx)
			return
		}

		c.Resp.RespSuccessWithData(
			gin.H{
				"data":   data,
				"count":  count,
				"cursor": cursor,
			}, ctx,
		)
		return
	}

	c.Resp.RespSuccessWithData(nil, ctx)
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

func descHashPageData(wait []string) ([]map[string]string, error) {
	if !isPageOk(len(wait)) {
		return nil, errors.New("数据异常！")
	}
	res, last := make([]map[string]string, 0), make([]map[string]string, 0)
	tempKey, temVal := -1, ""
	for i, val := range wait {
		if tempKey > -1 {
			res = append(
				res, map[string]string{
					temVal: val,
				},
			)
			tempKey = -1
			continue
		}
		temVal = val
		tempKey = i
	}
	//二次处理
	for _, temp := range res {
		for key, val := range temp {
			last = append(
				last, map[string]string{
					"key":   key,
					"value": val,
				},
			)
		}
	}
	return last, nil
}

func isPageOk(num int) bool {
	intRes := num / 2
	floatRes := float64(num) / 2.0
	if float64(intRes) == floatRes {
		return true
	} else {
		return false
	}
}

func getLenByKey(rd *redis.Client, key, keyType string) string {
	lenMsg := ""
	switch keyType {
	case "hash":
		lenMsg = fmt.Sprintf("%d", rd.HLen(key).Val())
		break
	case "list":
		lenMsg = fmt.Sprintf("%d", rd.LLen(key).Val())
		break
	case "set":
		lenMsg = fmt.Sprintf("%d", rd.SCard(key).Val())
		break
	case "zset":
		lenMsg = fmt.Sprintf("%d", rd.ZCard(key).Val())
		break
	case "string":
		lenMsg = fmt.Sprintf("%d", rd.StrLen(key).Val())
		break

	}
	return lenMsg
}
