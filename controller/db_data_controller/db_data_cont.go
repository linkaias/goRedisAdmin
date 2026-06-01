package db_data_controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"goRedisAdmin/controller"
	"goRedisAdmin/global/global_redis"
	"goRedisAdmin/utils/exoprt_utils"
	"goRedisAdmin/utils/log_utils"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

// dbDataCont implements Redis DB/key/value HTTP handlers.
type dbDataCont struct {
	controller.BaseController
}

// NewDbDataController creates a DB data controller with base init.
func NewDbDataController() DbDataController {
	cont := new(dbDataCont)
	cont.BaseInit()
	return cont
}

// DbDataController defines all Redis data management endpoints.
type DbDataController interface {
	DbList(ctx *gin.Context)
	GetKeys(ctx *gin.Context)
	GetVal(ctx *gin.Context)
	DelKey(ctx *gin.Context)
	AddVal(ctx *gin.Context)
	Flush(ctx *gin.Context)
	GetValByKey(ctx *gin.Context)
	// ExpireKey updates key TTL.
	ExpireKey(ctx *gin.Context)
	ExportKey(ctx *gin.Context)
}

// DbList returns all supported Redis DB indices with key counts.
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

// getDbKeyLen returns number of keys in the specified Redis DB.
func getDbKeyLen(i int) (uint, error) {
	rd, err := global_redis.GetRedisClient(i)
	if err != nil {
		return 0, err
	}
	defer rd.Close()
	num := rd.DbSize().Val()
	return uint(num), nil
}

// GetKeys lists keys in a selected DB, including type, length and expire info.
func (c dbDataCont) GetKeys(ctx *gin.Context) {
	// db_num defaults to 0 when parsing fails.
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
	// Build frontend table rows.
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
			temp["expire_at"] = expire.String() + "后过期"
		}
		data = append(data, temp)
	}
	c.Resp.RespSuccessWithData(data, ctx)
}

// GetVal fetches value details using unified payload and type dispatcher.
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

// Flush clears selected DB or all DBs depending on query param `type`.
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

// DelKey deletes one or multiple keys separated by comma.
func (c dbDataCont) DelKey(ctx *gin.Context) {
	// db_num defaults to 0 when parsing fails.
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

// ExportKey exports selected keys to a JSON file and returns it as attachment.
func (c dbDataCont) ExportKey(ctx *gin.Context) {
	// db_num defaults to 0 when parsing fails.
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
	utils := new(exoprt_utils.ExportUtils)
	utils.LoadKeysData(rd, waitDelKey)
	filePath, err := utils.SaveFile()
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	// Configure file-download headers.
	ctx.Header("Content-Type", "application/json")
	ctx.Header("Content-Disposition", `attachment; filename="keys.json"`)
	ctx.Header("Content-Transfer-Encoding", "binary")
	// Read generated JSON file content.
	data, err := os.ReadFile(filePath)
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	// Send file content in response body.
	ctx.Data(http.StatusOK, "application/json", data)
}

// ExpireKey updates key expiration time in seconds.
// expire=0 means remove expiration and persist the key.
func (c dbDataCont) ExpireKey(ctx *gin.Context) {
	// db_num defaults to 0 when parsing fails.
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
	if expireInt == 0 { // Persist key (no expiration).
		rd.Persist(key)
	} else { // Set expiration.
		_, err = rd.Expire(key, time.Duration(expireInt)*time.Second).Result()
	}
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	c.Resp.RespSuccess(ctx)
}

// AddVal creates a key/value according to provided Redis type.
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

// GetValByKey fetches value by key with type-specific retrieval strategy.
func (c dbDataCont) GetValByKey(ctx *gin.Context) {
	// db_num defaults to 0 when parsing fails.
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
		// cursor defaults to 0 for HSCAN pagination.
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
	if s.DType == "set" {
		result, err := rd.SMembers(s.Key).Result()
		if err != nil {
			log_utils.WriteLog("err", err, nil)
			c.Resp.RespError(err.Error(), ctx)
			return
		}
		c.Resp.RespSuccessWithData(
			gin.H{
				"data":  result,
				"count": len(result),
			}, ctx,
		)
		return

	}
	if s.DType == "list" {
		result, err := rd.LRange(s.Key, 0, -1).Result()
		if err != nil {
			log_utils.WriteLog("err", err, nil)
			c.Resp.RespError(err.Error(), ctx)
			return
		}
		c.Resp.RespSuccessWithData(
			gin.H{
				"data":  result,
				"count": len(result),
			}, ctx,
		)
		return
	}
	if s.DType == "zset" {
		result, err := rd.ZRange(s.Key, 0, -1).Result()
		if err != nil {
			log_utils.WriteLog("err", err, nil)
			c.Resp.RespError(err.Error(), ctx)
			return
		}
		c.Resp.RespSuccessWithData(
			gin.H{
				"data":  result,
				"count": len(result),
			}, ctx,
		)
		return
	}
	if s.DType == "stream" {
		result, err := rd.XRange(s.Key, "-", "+").Result()
		if err != nil {
			log_utils.WriteLog("err", err, nil)
			c.Resp.RespError(err.Error(), ctx)
			return
		}
		data := make([]map[string]string, 0, len(result))
		for _, item := range result {
			data = append(
				data,
				map[string]string{
					"id":     item.ID,
					"fields": marshalStreamFields(item.Values),
				},
			)
		}
		c.Resp.RespSuccessWithData(
			gin.H{
				"data":  data,
				"count": len(data),
			}, ctx,
		)
		return
	}

	c.Resp.RespSuccessWithData(
		gin.H{
			"data":        []interface{}{},
			"count":       0,
			"cursor":      0,
			"unsupported": true,
		},
		ctx,
	)
}

// handleAddVal dispatches add operation by Redis data type.
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
	case "stream":
		return cont.AddStream()
	}
	return errors.New("type not supported ! ")
}

// handleGetVal dispatches get operation by Redis data type.
func handleGetVal(valType string, cont *DbDataHelpCont) (interface{}, error) {
	switch valType {
	case "string":
		return cont.GetString()
	case "list":
		return cont.GetList()
	case "set":
		return cont.GetSet()
	case "zset":
		return cont.GetZSet()
	case "hash":
		return cont.GetHash()
	case "stream":
		return cont.GetStream()
	}
	return nil, errors.New("type not supported ! ")
}

// descHashPageData converts flat HSCAN field/value list into
// [{"key": field, "value": value}] format.
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
	// Transform intermediate map into stable key/value list objects.
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

// isPageOk checks whether array length is even (field/value pairs).
func isPageOk(num int) bool {
	intRes := num / 2
	floatRes := float64(num) / 2.0
	if float64(intRes) == floatRes {
		return true
	} else {
		return false
	}
}

// getLenByKey returns type-specific length/size text for a key.
func getLenByKey(rd *redis.Client, key, keyType string) string {
	lenMsg := ""
	switch keyType {
	case "hash":
		lenMsg = fmt.Sprintf("%d 个", rd.HLen(key).Val())
	case "list":
		lenMsg = fmt.Sprintf("%d 个", rd.LLen(key).Val())
	case "set":
		lenMsg = fmt.Sprintf("%d 个", rd.SCard(key).Val())
	case "zset":
		lenMsg = fmt.Sprintf("%d 个", rd.ZCard(key).Val())
	case "string":
		lenMsg = convertBytes(rd.StrLen(key).Val())
	case "stream":
		lenMsg = fmt.Sprintf("%d 条", rd.XLen(key).Val())

	}
	return lenMsg
}

// marshalStreamFields formats stream fields into JSON text for direct UI display.
func marshalStreamFields(fields map[string]interface{}) string {
	if len(fields) == 0 {
		return "{}"
	}
	b, err := json.Marshal(fields)
	if err != nil {
		return "{}"
	}
	return string(b)
}

// convertBytes converts byte count into human-readable binary unit string.
func convertBytes(bytes int64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB"}
	var index int
	size := float64(bytes)

	for size >= 1024 && index < len(units)-1 {
		size /= 1024
		index++
	}

	return fmt.Sprintf("%.2f %s", size, units[index])
}
