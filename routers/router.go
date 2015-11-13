package routers

import (
	"github.com/astaxie/beego"
	"star/controllers"
)

func init() {
	//首页
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/first", &controllers.FirstController{})
	beego.Router("/second", &controllers.SecondController{})

	// 能力地图、新员工培训、新员工交流
	beego.Router("/capability", &controllers.CapabilityController{})
	beego.Router("/training", &controllers.TrainingController{})
	beego.Router("/communicate", &controllers.CommunicateController{})

	// 登录与注册
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/regist", &controllers.RegistController{})

}
