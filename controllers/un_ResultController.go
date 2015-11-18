package controllers

import (
	"github.com/astaxie/beego"
	"star/models"
)

type ResultController struct {
	beego.Controller
}

func (t *ResultController) Get() {
	t.Data["IsResult"] = true
	t.Data["IsLogin"] = CheckAccount(t.Ctx)
	namestruct, err := t.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		t.Data["UserName"] = name
	}
	t.Data["AllUser"], err = models.GetAllUser()
	if err != nil {
		beego.Error(err)
	}
	t.TplNames = "Result.html"
}
