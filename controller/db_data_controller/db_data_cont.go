package db_data_controller

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/controller"
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
}

func (c dbDataCont) DbList(ctx *gin.Context) {

}
