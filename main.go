package main

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// "star/controllers"
	"star/g"
	_ "star/routers"
)

func init() {
	// models.RegistDB()
}
func main() {
	// orm.Debug = true
	// models.RegistDB()
	// orm.RunSyncdb("default", false, true)
	//models.AddUser()
	//models.AddTotalAndWin("martin", true)
	g.InitEnv()
	//go controllers.GetCapabilityFromMysqlPerHour()//减少查询数据库
	//models.AddCapabilitiesData()
	//models.AddCapabilityMap()
	beego.Run()
}
