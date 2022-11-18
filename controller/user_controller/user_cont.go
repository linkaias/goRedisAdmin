package user_controller

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/controller"
)

type userController struct {
	controller.BaseController
}

func NewUserController() UserController {
	cont := &userController{}
	cont.BaseInit()
	return cont
}

type UserController interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

func (c userController) Login(ctx *gin.Context) {

}

func (c userController) Logout(ctx *gin.Context) {

}
