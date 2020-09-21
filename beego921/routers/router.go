package routers

import (
	"beego921/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
