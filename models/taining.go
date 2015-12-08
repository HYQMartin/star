package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type TrainingSchedulePublish struct {
	Id          int64
	Title       string
	Teacher     string
	Content     string
	Description string
	StartDate   time.Time
	Room        string
	Status      string
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
}

type TrainingScheduleCollect struct {
	Id             int64
	Title          string
	Description    string
	Understand     string
	ExpectDate     string
	ImportantLevel string
	Fields         string
	Created        time.Time `orm:"auto_now_add;type(datetime)"`
}

func AddTrainingSchedulePublishs() {

	AddTrainingSchedulePublish("linux 基础", "董兴水D000", "linux,bash,kernel", "", "H2-4A22R", "未开课", time.Now())
	AddTrainingSchedulePublish("makefile 基础", "雷歆L000", "makefile", "", "H2-4A22R", "未开课", time.Now())
	AddTrainingSchedulePublish("C 基础", "刘剑青L202", "C 基础", "", "H2-4A22R", "未开课", time.Now())
}

func AddTrainingSchedulePublish(title, teacher, content, description, room, status string, startdate time.Time) string {
	o := orm.NewOrm()

	//验证合法性
	qs := o.QueryTable("training_schedule_publish")

	var c TrainingSchedulePublish

	err := qs.Filter("Title", title).One(&c)
	if err != orm.ErrNoRows {
		return "exist such Title"
	}

	t := new(TrainingSchedulePublish)
	t.Title = title
	t.Teacher = teacher
	t.Content = content
	t.Description = description
	t.StartDate = startdate
	t.Room = room
	t.Status = status
	_, err = o.Insert(t)
	if err != nil {
		return err.Error()
	}

	return "success"
}

func AddTrainingScheduleCollects() {

	AddTrainingScheduleCollect("申请课程 linux 基础", "linux,bash,kernel", "不了解", "这个月", "紧急", "docker")
}

func AddTrainingScheduleCollect(title, description, understand, expectdate, important, fields string) string {
	o := orm.NewOrm()

	//验证合法性
	qs := o.QueryTable("training_schedule_collect")

	var c TrainingScheduleCollect

	err := qs.Filter("Title", title).One(&c)
	if err != orm.ErrNoRows {
		return "exist such Title"
	}

	t := new(TrainingScheduleCollect)
	t.Title = title
	t.Description = description
	t.Understand = understand
	t.ExpectDate = expectdate
	t.ImportantLevel = important
	t.Fields = fields

	_, err = o.Insert(t)
	if err != nil {
		return err.Error()
	}

	return "success"
}

func ALLTrainingSchedulePublish() (all []TrainingSchedulePublish, err error) {
	all = make([]TrainingSchedulePublish, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("training_schedule_publish")
	_, err = qs.All(&all)
	return
}

func ALLTrainingScheduleCollect() (all []TrainingScheduleCollect, err error) {
	all = make([]TrainingScheduleCollect, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("training_schedule_collect")
	_, err = qs.All(&all)
	return
}
