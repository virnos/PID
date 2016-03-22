package routers

import (
	"PID/controllers"
	cms "PID/controllers/cms"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/cms/device", &cms.ProjectController{})
}
