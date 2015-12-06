package controllers

import (
	"star/g"
	"star/models/blog"
	"star/models/catalog"
)

type MainController struct {
	BaseController
}

func (this *MainController) Get() {
	this.Data["Catalogs"] = catalog.All()
	this.Data["PageTitle"] = "首页"
	this.Data["Addblog"] = ""
	this.Layout = "layout/default.html"
	this.TplNames = "index.html"
}

func (this *MainController) Read() {
	ident := this.GetString(":ident")
	b := blog.OneByIdent(ident)
	if b == nil {
		this.Ctx.WriteString("no such article")
		return
	}

	b.Views = b.Views + 1
	blog.Update(b, "")

	this.Data["Blog"] = b
	this.Data["Content"] = g.RenderMarkdown(blog.ReadBlogContent(b).Content)
	this.Data["PageTitle"] = b.Title
	this.Data["Catalog"] = catalog.OneById(b.CatalogId)
	this.Data["Addblog"] = `<a href="/catalog/share">返回</a>`
	this.Layout = "layout/default.html"
	this.TplNames = "article/read.html"
}

func (this *MainController) ListByCatalog() {
	cata := this.Ctx.Input.Param(":ident")
	if cata == "" {
		this.Ctx.WriteString("catalog ident is blank")
		return
	}
	if cata == "capability" {
		this.Redirect("/capability", 302)
	}
	if cata == "training" {
		this.Redirect("/training", 302)
	}
	if cata == "newemployee" {
		this.Redirect("/newemployee", 302)
	}
	limit := this.GetIntWithDefault("limit", 10)

	c := catalog.OneByIdent(cata)
	if c == nil {
		this.Ctx.WriteString("catalog:" + cata + " not found")
		return
	}

	ids := blog.Ids(c.Id)
	pager := this.SetPaginator(limit, int64(len(ids)))
	blogs := blog.ByCatalog(c.Id, pager.Offset(), limit)

	this.Data["Catalog"] = c
	this.Data["Blogs"] = blogs
	this.Data["PageTitle"] = c.Name
	this.Data["Addblog"] = `<a href="/admin/article/add">新增总结分享</a><div><a href="/">  返回</a></div>`
	this.Layout = "layout/default.html"

	this.TplNames = "article/by_catalog.html"
}
