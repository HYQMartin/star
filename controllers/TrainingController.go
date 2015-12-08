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
	c, err := models.ALLTrainingScheduleCollect()
	if err != nil {
		return
	}

	// c, err := models.ALLTrainingScheduleCollect()
	// if err != nil {
	// 	return
	// }
	t.Data["ScheduleTable"] = p
	t.Data["CollectTable"] = c

	return
}

type TrainingPublishController struct {
	BaseController
}

func (t *TrainingPublishController) Post() {
	if !t.IsLogin && !t.IsAdmin {
		t.Ctx.WriteString("you are not login or admin")
		return
	}
	name := t.GetString("name")
	content := t.GetString("content")
	descp := t.GetString("descp")
	tt := t.GetString("time")
	place := t.GetString("place")
	teacher := t.GetString("teacher")
	status := t.GetString("status")
	tm2, err := time.Parse("2006-01-02 15:04:05", tt)
	if err != nil {
		t.Ctx.WriteString(err.Error())
		fmt.Print("Parse", err)
		return
	}
	models.AddTrainingSchedulePublish(name, teacher, content, descp, place, status, tm2)
	t.Ctx.WriteString("success")
	return
}

type TrainingCollectController struct {
	BaseController
}

func (t *TrainingCollectController) Post() {
	if !t.IsLogin && !t.IsAdmin {
		t.Ctx.WriteString("you are not login or admin")
		return
	}
	title := t.GetString("title")
	desc := t.GetString("desc")
	level := t.GetString("level")
	expecttime := t.GetString("expecttime")
	important := t.GetString("important")

	if title != "" && desc != "" && level != "" && expecttime != "" && important != "" {
		ret := models.AddTrainingScheduleCollect(title, desc, level, expecttime, important, "")
		t.Ctx.WriteString(ret)
		return
	}

	t.Ctx.WriteString("faild")
	return
}
