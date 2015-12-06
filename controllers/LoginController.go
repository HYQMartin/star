package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"star/g"
	"star/models"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {
	this.TplNames = "login&regist/login.html"
}

func (this *LoginController) DoLogin() {
	name := this.GetString("name")
	if name == "" {
		this.Ctx.WriteString("name is blank")
		return
	}
	password := this.GetString("password")
	if password == "" {
		this.Ctx.WriteString("password is blank")
		return
	}

	if g.RootName == name && g.RootPass == password {
		this.Ctx.SetCookie("admin_name", name, 2592000, "/")
		this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "admin_password="+password+"; Max-Age=2592000; Path=/; httponly")
		this.Ctx.WriteString("admin login success")
		return
	}
	result := models.CheckPassword(name, password)
	if result {
		this.Ctx.SetCookie("name", name, 2592000, "/")
		this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "password="+password+"; Max-Age=2592000; Path=/; httponly")
		this.Ctx.WriteString("login success")
		return
	}

	return
}

func (this *LoginController) Logout() {
	this.Ctx.SetCookie("admin_name", g.RootName, 0, "/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "admin_password="+g.RootPass+"; Max-Age=0; Path=/; httponly")
	this.Ctx.SetCookie("name", g.RootName, 0, "/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "password="+g.RootPass+"; Max-Age=0; Path=/; httponly")
	this.Redirect("/", 302)

	return
	fmt.Println("------------------")
}

func CheckAccount(t *context.Context) bool {
	n, err := t.Request.Cookie("name")
	if err != nil {
		return false
	}
	name := n.Value

	p, err := t.Request.Cookie("password")
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

func (this *RegistController) Post() {
	name := this.GetString("name")
	if name == "" {
		this.Ctx.WriteString("name is blank")
		return
	}
	password := this.GetString("password")
	if password == "" {
		this.Ctx.WriteString("password is blank")
		return
	}
	result := models.RegistUser(name, password)

	this.Ctx.WriteString(result)
	return
}
