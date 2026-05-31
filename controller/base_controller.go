package controller

import (
	"goRedisAdmin/global/global_response"
	"goRedisAdmin/utils/log_utils"
	"io"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

// BaseController provides shared helper methods for all controllers.
// It wraps response builder and request parameter conversion helpers.
type BaseController struct {
	Resp *global_response.Response
}

// BaseInit initializes the response helper instance.
func (b *BaseController) BaseInit() {
	b.Resp = new(global_response.Response)
}

// JSONToStruct unmarshals a JSON string into the provided target object.
// Any parsing error is logged and returned.
func (b *BaseController) JSONToStruct(jsonStr string, obj interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), obj)
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		return err
	}
	return nil
}

// GetBodyByRequest reads request body bytes based on Content-Length.
//
// Note: This helper performs a single read and is designed for current
// request patterns in this codebase.
func (b *BaseController) GetBodyByRequest(ctx *gin.Context) ([]byte, error) {
	body := make([]byte, ctx.Request.ContentLength)
	_, err := ctx.Request.Body.Read(body)
	if err != nil && err != io.EOF {
		log_utils.WriteLog("err", err, nil)
		return nil, err
	}
	return body, nil
}

// ParamToInt reads a GET/POST parameter and parses it as int.
//
// method accepts "get"/"GET" or "post"/"POST".
func (b *BaseController) ParamToInt(ctx *gin.Context, param string, method string) (int, error) {
	intStr := ""
	if method == "get" || method == "GET" {
		intStr = ctx.Query(param)
	} else if method == "post" || method == "POST" {
		intStr = ctx.PostForm(param)
	}
	return strconv.Atoi(intStr)
}

// ParamToInt64 reads a GET/POST parameter and parses it as int64.
//
// method accepts "get"/"GET" or "post"/"POST".
func (b *BaseController) ParamToInt64(ctx *gin.Context, param string, method string) (int64, error) {
	intStr := ""
	if method == "get" || method == "GET" {
		intStr = ctx.Query(param)
	} else if method == "post" || method == "POST" {
		intStr = ctx.PostForm(param)
	}
	res, err := strconv.ParseInt(intStr, 10, 64)
	return res, err
}
