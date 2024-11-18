package info_controller

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/controller"
	"goRedisAdmin/global/global_redis"
	"goRedisAdmin/utils/log_utils"
	"strings"
)

type infoController struct {
	controller.BaseController
}

func NewInfoController() InfoController {
	cont := &infoController{}
	cont.BaseInit()
	return cont
}

type InfoController interface {
	GetInfo(ctx *gin.Context)
}

func (c infoController) GetInfo(ctx *gin.Context) {
	rd, err := global_redis.GetRedisClient(0)
	if err != nil {
		log_utils.WriteLog("err", err, nil)
		c.Resp.RespError(err.Error(), ctx)
		return
	}
	defer rd.Close()
	//ctxSon := context.Background()
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
