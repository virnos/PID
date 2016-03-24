package routers

import (
	"PID/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/project", &controllers.ProjectController{})
}
