package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/ulricqin/goutils/filetool"
	"github.com/ulricqin/goutils/paginator"
	"github.com/ulricqin/goutils/strtool"
	"star/g"
	"star/models"
	"star/models/catalog"
	"strconv"
	"strings"
	"time"
)

const (
	EDITOR_IMG_DIR = "static/uploads/editor"
)

type Checker interface {
	CheckLogin()
}

type BaseController struct {
	beego.Controller
	IsAdmin  bool
	IsLogin  bool
	UserName string
}

func (this *BaseController) Prepare() {
	this.Data["BlogLogo"] = g.BlogLogo
	this.Data["BlogTitle"] = g.BlogTitle
	this.Data["BlogResume"] = g.BlogResume
	this.Data["RootName"] = g.RootName
	this.Data["RootEmail"] = g.RootEmail
	this.Data["RootPortrait"] = g.RootPortrait
	this.AssignIsAdmin()
	if app, ok := this.AppController.(Checker); ok {
		app.CheckLogin()
	}
}

func (this *BaseController) AssignIsAdmin() {
	admin_name := this.Ctx.GetCookie("admin_name")
	admin_password := this.Ctx.GetCookie("admin_password")
	name := this.Ctx.GetCookie("name")
	password := this.Ctx.GetCookie("password")

	if admin_name == g.RootName && admin_password == g.RootPass && admin_name != "" && admin_password != "" {
		this.IsAdmin = true
		this.IsLogin = true
		this.UserName = "管理员"
		// fmt.Println("*******************登录成功，username=", this.UserName)
	} else {
		this.IsAdmin = false
		if models.CheckPassword(name, password) {
			this.IsLogin = true
			this.UserName = name
		} else {
			this.IsLogin = false
		}
	}

	this.Data["IsAdmin"] = this.IsAdmin
	this.Data["IsLogin"] = this.IsLogin
	this.Data["UserName"] = this.UserName
}

func (this *BaseController) SetPaginator(per int, nums int64) *paginator.Paginator {
	p := paginator.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}

func (this *BaseController) GetIntWithDefault(paramKey string, defaultVal int) int {
	valStr := this.GetString(paramKey)
	var val int
	if valStr == "" {
		val = defaultVal
	} else {
		var err error
		val, err = strconv.Atoi(valStr)
		if err != nil {
			val = defaultVal
		}
	}
	return val
}

func (this *BaseController) JsStorage(action, key string, values ...string) {
	value := action + ":::" + key
	if len(values) > 0 {
		value += ":::" + values[0]
	}
	this.Ctx.SetCookie("JsStorage", value, 1<<31-1, "/")
}

//************************************************************************
type ApiController struct {
	BaseController
}

func (this *ApiController) Health() {
	fmt.Println(catalog.All()[0])
	this.Ctx.WriteString("ok")
}

func (this *ApiController) Md5() {
	p := this.GetString("p")
	this.Ctx.WriteString(strtool.Md5(strings.TrimSpace(p)))
}

func (this *ApiController) Upload() {
	result := map[string]interface{}{
		"success": false,
	}

	defer func() {
		this.Data["json"] = &result
		this.ServeJson()
	}()

	_, header, err := this.GetFile("image")
	if err != nil {
		return
	}

	ext := filetool.Ext(header.Filename)
	imgPath := fmt.Sprintf("%s/%d%s", EDITOR_IMG_DIR, time.Now().Unix(), ext)

	filetool.InsureDir(EDITOR_IMG_DIR)
	err = this.SaveToFile("image", imgPath)
	if err != nil {
		return
	}

	imgUrl := "/" + imgPath
	if g.UseQiniu {
		if addr, er := g.UploadFile(imgPath, imgPath); er != nil {
			return
		} else {
			imgUrl = addr
			filetool.Unlink(imgPath)
		}
	}

	result["link"] = imgUrl
	result["success"] = true

}

func (this *ApiController) Markdown() {
	if this.IsAjax() {
		result := map[string]interface{}{
			"success": false,
		}
		action := this.GetString("action")
		switch action {
		case "preview":
			content := this.GetString("content")
			result["preview"] = g.RenderMarkdown(content)
			result["success"] = true
		}
		this.Data["json"] = result
		this.ServeJson()
	}
}
