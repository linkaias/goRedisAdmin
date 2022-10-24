package controller

import (
	"goRedisAdmin/global/global_response"
)

type BaseController struct {
	Resp *global_response.Response
}

func (b *BaseController) BaseInit() {
	b.Resp = new(global_response.Response)
}
