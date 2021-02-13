package routers

import (
	"github.com/astaxie/beego"
	"meetingRecord/controllers"
)

func init() {
	beego.Router("/default", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/exit", &controllers.ExitController{})
	beego.Router("/projects", &controllers.ProjectsController{})
	beego.Router("/new_text", &controllers.NewTextController{})
	beego.Router("/text/:id", &controllers.ShowTextController{})
	beego.Router("/text/update", &controllers.UpdateTextController{})
	beego.Router("/text/delete", &controllers.DeleteTextController{})
}
