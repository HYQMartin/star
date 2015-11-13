package controllers

import (
	"github.com/astaxie/beego"
)

type CapabilityController struct {
	beego.Controller
}

func (c *CapabilityController) Get() {
	c.TplNames = "Capability.html"
	return
}
