package main

import (
	"flag"
	"github.com/astaxie/beego"
	"short_url/controllers"
)

var (
	port string
)

func main() {
	flag.StringVar(&port, "port", ":8080", "port to listen")
	flag.Parse()
	beego.RESTRouter("/api/url", &controllers.BaseController{})
	beego.RESTRouter("/", &controllers.RedirectController{})
	beego.Run()
}
