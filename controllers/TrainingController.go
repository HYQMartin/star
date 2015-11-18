package controllers

import (
	"github.com/astaxie/beego"
)

type TrainingController struct {
	beego.Controller
}

func (t *TrainingController) Get() {
	t.TplNames = "main_training/Training.html"
	return
}
