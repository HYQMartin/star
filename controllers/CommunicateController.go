package controllers

import (
	// "github.com/astaxie/beego"
	// "star/g"
	"star/models/blog"
	"star/models/catalog"
)

type CommunicateController struct {
	BaseController
}

func (this *CommunicateController) Get() {
	this.Data["Catalogs"] = catalog.NAll()
	this.Data["PageTitle"] = "新员工培养"
	this.Data["Addblog"] = ""
	this.Layout = "main_newemployee/layout/default.html"
	this.TplNames = "main_newemployee/index.html"
}

func (this *CommunicateController) Read() {
	ident := this.GetString(":ident")
	b := blog.NOneByIdent(ident)
	if b == nil {
		this.Ctx.WriteString("no such article")
		return
	}

	b.Views = b.Views + 1
	blog.NUpdate(b, "")

	this.Data["Blog"] = b
	this.Data["Content"] = blog.NReadBlogContent(b).Content
	this.Data["PageTitle"] = b.Title
	tmp := catalog.NOneById(b.CatalogId)
	this.Data["Catalog"] = tmp
	this.Data["Addblog"] = `<a href="/newemployee/catalog/` + tmp.Ident + `">返回</a>`
	this.Layout = "main_newemployee/layout/default.html"
	this.TplNames = "main_newemployee/article/read.html"
}

func (this *CommunicateController) ListByCatalog() {
	cata := this.Ctx.Input.Param(":ident")
	if cata == "" {
		this.Ctx.WriteString("catalog ident is blank")
		return
	}
	limit := this.GetIntWithDefault("limit", 10)

	c := catalog.NOneByIdent(cata)
	if c == nil {
		this.Ctx.WriteString("catalog:" + cata + " not found")
		return
	}

	ids := blog.NIds(c.Id)
	pager := this.SetPaginator(limit, int64(len(ids)))
	blogs := blog.NByCatalog(c.Id, pager.Offset(), limit)

	this.Data["Catalog"] = c
	this.Data["Blogs"] = blogs
	this.Data["PageTitle"] = c.Name
	this.Data["Addblog"] = `<a href="/newemployee/admin/article/add?cident=` + c.Ident + `">新增帖子</a>  <div><a href="/newemployee">  返回</a></div>`
	this.Layout = "main_newemployee/layout/default.html"
	this.TplNames = "main_newemployee/article/by_catalog.html"
}
