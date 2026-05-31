package global_response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RespPageResult is the canonical paged payload shape returned by APIs.
type RespPageResult struct {
	List     []map[string]interface{} `json:"list"`
	Total    int64                    `json:"total"`
	Page     int                      `json:"page"`
	PageSize int                      `json:"pageSize"`
}

// Response is the unified API response envelope.
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	// ERROR indicates a generic business or request error.
	ERROR = 7
	// NoLOGIN indicates authentication failure or missing login token.
	NoLOGIN = 6
	// SUCCESS indicates successful request handling.
	SUCCESS = 0
)

// Response writes the standard response envelope as JSON with HTTP 200.
func (f Response) Response(code int, msg string, data interface{}, ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK, Response{
			code,
			msg,
			data,
		},
	)
}

// RespPage writes a successful paged payload.
func (f Response) RespPage(data *RespPageResult, ctx *gin.Context) {
	f.Response(SUCCESS, "Success !", data, ctx)
}

// RespSuccess writes a successful empty object response.
func (f Response) RespSuccess(ctx *gin.Context) {
	f.Response(SUCCESS, "Success !", map[string]interface{}{}, ctx)
}

// RespSuccessWithMsg writes a successful empty data response with custom message.
func (f Response) RespSuccessWithMsg(message string, ctx *gin.Context) {
	f.Response(SUCCESS, message, map[string]interface{}{}, ctx)
}

// RespSuccessWithData writes a successful response with custom payload.
func (f Response) RespSuccessWithData(data interface{}, ctx *gin.Context) {
	f.Response(SUCCESS, "Success !", data, ctx)
}

// RespError writes a standardized error response with custom message.
func (f Response) RespError(msg string, ctx *gin.Context) {
	f.Response(ERROR, msg, map[string]interface{}{}, ctx)
}
