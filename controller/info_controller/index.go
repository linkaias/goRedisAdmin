package info_controller

import (
	"goRedisAdmin/controller"
	"goRedisAdmin/global/global_redis"
	"goRedisAdmin/utils/log_utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// infoController implements Redis INFO-related handlers.
type infoController struct {
	controller.BaseController
}

// NewInfoController creates an info controller instance with base init.
func NewInfoController() InfoController {
	cont := &infoController{}
	cont.BaseInit()
	return cont
}

// InfoController defines info API handlers.
type InfoController interface {
	GetInfo(ctx *gin.Context)
}

// GetInfo fetches Redis INFO output and formats it for frontend display.
// Line breaks and tabs are converted into HTML-friendly strings.
func (c infoController) GetInfo(ctx *gin.Context) {
	rd, err := global_redis.GetRedisClient(0)
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer rd.Close()
	res, err := rd.Info().Result()
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	res = strings.ReplaceAll(res, "\n", "<br/>")
	res = strings.ReplaceAll(res, "\t", "&nbsp;&nbsp;&nbsp;&nbsp;")
	c.Resp.RespSuccessWithData(res, ctx)
}
