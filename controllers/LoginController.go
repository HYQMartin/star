package controllers

import (
	"fmt"
	"star/g"
)

var IsLogin = false

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

	if g.RootName != name {
		this.Ctx.WriteString("name is incorrect")
		return
	}

	fmt.Println("password=", password)
	if g.RootPass != password {
		this.Ctx.WriteString("password is incorrect")
		return
	}

	IsLogin = true
	this.Ctx.SetCookie("bb_name", name, 2592000, "/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "bb_password="+password+"; Max-Age=2592000; Path=/; httponly")

	this.Ctx.WriteString("")
}

func (this *LoginController) Logout() {
	this.Ctx.SetCookie("bb_name", g.RootName, 0, "/")
	this.Ctx.ResponseWriter.Header().Add("Set-Cookie", "bb_password="+g.RootPass+"; Max-Age=0; Path=/; httponly")
	this.Ctx.WriteString("logout")
	IsLogin = false
	this.Redirect("/", 302)
	fmt.Println("------------------")
}
