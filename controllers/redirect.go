package controllers

type RedirectController struct {
	BaseController
}

//redirect to /api/url
func (this *RedirectController) Get() {
	if this.Ctx.Request.URL.Path == "/" {
		this.Data["title_name"] = "短链接服务_beego"
		this.TplName = "index.html"
		return
	}
	this.Redirect("/api/url", 302)
}
