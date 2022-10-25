package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"goRedisAdmin/global/global_response"
	"io"
	"strconv"
)

type BaseController struct {
	Resp *global_response.Response
}

func (b *BaseController) BaseInit() {
	b.Resp = new(global_response.Response)
}

//JSONToStruct 将json转换为结构体
func (b *BaseController) JSONToStruct(jsonStr string, obj interface{}) error {
	err := json.Unmarshal([]byte(jsonStr), obj)
	if err != nil {
		return err
	}
	return nil
}

//GetBodyByRequest 获取请求体body中的数据
func (b *BaseController) GetBodyByRequest(ctx *gin.Context) ([]byte, error) {
	body := make([]byte, ctx.Request.ContentLength)
	_, err := ctx.Request.Body.Read(body)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return body, nil
}

// ParamToInt 把请求参数转化为int
//method "get"/"post"
func (b *BaseController) ParamToInt(ctx *gin.Context, param string, method string) (int, error) {
	intStr := ""
	if method == "get" || method == "GET" {
		intStr = ctx.Query(param)
	} else if method == "post" || method == "POST" {
		intStr = ctx.PostForm(param)
	}
	return strconv.Atoi(intStr)
}

// ParamToInt64 把请求参数转化为int64
//method "get"/"post"
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
