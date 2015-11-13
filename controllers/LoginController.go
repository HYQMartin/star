package controllers

import (
	"beego/wuziqi/models"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	//"time"
)

type LoginController struct {
	beego.Controller
}

func (t *LoginController) Get() {
	// 判断是否为退出操作
	if t.Input().Get("exit") == "true" {
		t.Ctx.SetCookie("name", "", -1, "/")
		t.Ctx.SetCookie("pwd", "", -1, "/")
		t.Redirect("/", 302)
		return
	}
	t.Data["IsLogin"] = CheckAccount(t.Ctx)
	t.TplNames = "Login.html"
	return
}
func (t *LoginController) Post() {
	name := t.Input().Get("uname")
	pwd := t.Input().Get("pwd")
	autoLogin := t.Input().Get("autoLogin")
	fmt.Println("name", name, "pwd", pwd, "auto", autoLogin)
	isright := models.CheckPassword(name, pwd)
	fmt.Println("5555555555", "name", name, "pwd", pwd, isright)

	maxAge := 0
	maxAge = 1<<31 - 1

	t.Ctx.SetCookie("name", name, maxAge, "/")
	t.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	if isright {
		t.TplNames = "Home.html"
		/*t.Data["Author"] = "叶青，子龙"
		t.Data["Time"] = time.Now().Format(time.Stamp)
		t.Data["IsHome"] = true
		t.Data["IsLogin"] = true*/
		t.Redirect("/", 302)
		return
	}
	namestruct, err := t.Ctx.Request.Cookie("name")
	if err == nil {
		aname := namestruct.Value
		t.Data["UserName"] = aname
	}

	t.TplNames = "Login.html"
	return
}
func CheckAccount(t *context.Context) bool {
	n, err := t.Request.Cookie("name")
	if err != nil {
		return false
	}
	name := n.Value

	p, err := t.Request.Cookie("pwd")
	if err != nil {
		return false
	}
	password := p.Value
	isright := models.CheckPassword(name, password)
	return isright
}

type RegistController struct {
	beego.Controller
}

func (t *RegistController) Get() {
	//t.Data["IsResult"] = true
	t.Data["IsLogin"] = CheckAccount(t.Ctx)
	namestruct, err := t.Ctx.Request.Cookie("name")
	if err == nil {
		name := namestruct.Value
		t.Data["UserName"] = name
	}

	t.TplNames = "Regist.html"
	return
}
func (t *RegistController) Post() {
	name := t.Input().Get("name")
	password := t.Input().Get("password")
	phone := t.Input().Get("phone")
	result := models.RegistUser(name, password, phone)

	var re map[string]string
	re = make(map[string]string)
	re["result"] = result
	res, _ := json.Marshal(re)
	t.Ctx.WriteString(string(res))

	return
}
