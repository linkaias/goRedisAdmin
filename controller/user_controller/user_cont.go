package user_controller

import (
	"github.com/gin-gonic/gin"
	"goRedisAdmin/controller"
	"goRedisAdmin/global/initData"
	"goRedisAdmin/utils"
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

type LoginStruct struct {
	User string `json:"user" form:"user"`
	Pwd  string `json:"pwd" form:"pwd"`
}

func (c userController) Login(ctx *gin.Context) {
	info := new(LoginStruct)
	_ = ctx.ShouldBind(info)
	//数据库查询
	cfg := initData.IniRead.Section("admin")
	user := cfg.Key("username").String()
	passwd := cfg.Key("password").String()
	if user != info.User || passwd != info.Pwd {
		c.Resp.RespError("用户名或密码错误！", ctx)
		return
	}
	//登录成功 生成token
	token, _ := utils.BCYGenerateToken(info.User)

	c.Resp.RespSuccessWithData(token, ctx)
}

func (c userController) Logout(ctx *gin.Context) {

}
