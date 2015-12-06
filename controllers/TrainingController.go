package controllers

import (
	"fmt"
	// "github.com/astaxie/beego"
	"star/models"
	"time"
)

type TrainingController struct {
	BaseController
}

func (t *TrainingController) Get() {
	t.TplNames = "main_training/Training.html"

	p, err := models.ALLTrainingSchedulePublish()
	if err != nil {
		return
	}
	// c, err := models.ALLTrainingScheduleCollect()
	// if err != nil {
	// 	return
	// }
	t.Data["ScheduleTable"] = p

	return
}

type TrainingPublishController struct {
	BaseController
}

func (t *TrainingPublishController) Post() {
	name := t.GetString("name")
	content := t.GetString("content")
	descp := t.GetString("descp")
	tt := t.GetString("time")
	place := t.GetString("place")
	teacher := t.GetString("teacher")
	status := t.GetString("status")
	fmt.Print("*********time********", tt)
	tm2, err := time.Parse("2006-01-02 15:04:05", tt)
	if err != nil {
		t.Ctx.WriteString("error")
		fmt.Print("Parse", err)
		return
	}
	models.AddTrainingSchedulePublish(name, teacher, content, descp, place, status, tm2)
	t.Ctx.WriteString("success")
	return
}
