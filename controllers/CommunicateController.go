package controllers

import (
	"github.com/astaxie/beego"
)

type CommunicateController struct {
	beego.Controller
}

func (c *CommunicateController) Get() {
	c.TplNames = "Communicate.html"
	return
}
