package user_controller

import (
	"goRedisAdmin/controller"
	"goRedisAdmin/global/initData"
	"goRedisAdmin/utils"

	"github.com/gin-gonic/gin"
)

// userController implements authentication handlers.
type userController struct {
	controller.BaseController
}

// NewUserController creates a user controller with initialized base response.
func NewUserController() UserController {
	cont := &userController{}
	cont.BaseInit()
	return cont
}

// UserController declares login/logout endpoints.
type UserController interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

// LoginStruct represents login request payload.
type LoginStruct struct {
	User string `json:"user" form:"user"`
	Pwd  string `json:"pwd" form:"pwd"`
}

// Login validates provided credentials against config.ini admin section.
// If valid, it issues a JWT token and returns it to the client.
func (c userController) Login(ctx *gin.Context) {
	info := new(LoginStruct)
	_ = ctx.ShouldBind(info)
	// Compare with configured admin credentials.
	cfg := initData.IniRead.Section("admin")
	user := cfg.Key("username").String()
	passwd := cfg.Key("password").String()
	if user != info.User || passwd != info.Pwd {
		c.Resp.RespError("用户名或密码错误！", ctx)
		return
	}
	// Issue token on successful login.
	token, _ := utils.BCYGenerateToken(info.User)

	c.Resp.RespSuccessWithData(token, ctx)
}

// Logout is currently a placeholder endpoint.
func (c userController) Logout(ctx *gin.Context) {

}
