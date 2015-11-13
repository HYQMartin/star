package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"star/models"
	_ "star/routers"
)

func init() {
	models.RegistDB()
}
func main() {
	orm.Debug = true
	// models.RegistDB()
	orm.RunSyncdb("default", false, true)
	//models.AddUser()
	//models.AddTotalAndWin("martin", true)
	beego.Run()
}
