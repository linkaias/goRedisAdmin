package global_response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RespPageResult struct {
	List     []map[string]interface{} `json:"list"`
	Total    int64                    `json:"total"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"pageSize"`
}

// Response 全局响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

// Response 自定义返回数据
func (f Response) Response(code int, msg string, data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

// RespPage 返回分页数据格式
func (f Response) RespPage(data *RespPageResult, ctx *gin.Context) {
	f.Response(SUCCESS, "Success !", data, ctx)
}

// RespSuccess 操作成功
func (f Response) RespSuccess(ctx *gin.Context) {
	f.Response(SUCCESS, "Success !", map[string]interface{}{}, ctx)
}

// RespSuccessWithMsg 操作成功并自定义msg
func (f Response) RespSuccessWithMsg(message string, ctx *gin.Context) {
	f.Response(SUCCESS, message, map[string]interface{}{}, ctx)
}

// RespSuccessWithData 操作成功并自定义data
func (f Response) RespSuccessWithData(data interface{}, ctx *gin.Context) {
	f.Response(SUCCESS, "Success !", data, ctx)
}

// RespError 请求失败自定义错误信息
func (f Response) RespError(msg string, ctx *gin.Context) {
	f.Response(ERROR, msg, map[string]interface{}{}, ctx)
}
