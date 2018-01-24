package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) output(res string) {
	this.Ctx.WriteString(res)
}
func (this *BaseController) success(res interface{}) {
	this.Data["json"] = map[string]interface{}{
		"success": true,
		"result":  res,
	}
	this.ServeJSON()
}
