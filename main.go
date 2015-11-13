package main

import (
	"beego/wuziqi/models"
	_ "beego/wuziqi/routers"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
)

func init() {
	models.RegistDB()
}
func main() {
	//orm.Debug = true
	//models.RegistDB()
	//orm.RunSyncdb("default", false, true)
	//models.AddUser()
	//models.AddTotalAndWin("martin", true)
	beego.Run()
}
