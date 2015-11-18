package controllers

import (
	// "github.com/astaxie/beego"
	"star/g"
	"star/models/blog"
	"star/models/catalog"
)

type CommunicateController struct {
	BaseController
}

func (this *CommunicateController) Get() {
	this.Data["Catalogs"] = catalog.NAll()
	this.Data["PageTitle"] = "新员工培养"
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
	this.Data["Content"] = g.RenderMarkdown(blog.NReadBlogContent(b).Content)
	this.Data["PageTitle"] = b.Title
	this.Data["Catalog"] = catalog.NOneById(b.CatalogId)
	this.Layout = "main_newemployee/layout/default.html"
	this.TplNames = "main_newemployee/article/read.html"
}

func (this *CommunicateController) ListByCatalog() {
	cata := this.Ctx.Input.Param(":ident")
	if cata == "" {
		this.Ctx.WriteString("catalog ident is blank")
		return
	}
	// if cata == "capability" {
	// 	this.Redirect("/capability", 302)
	// }
	// if cata == "training" {
	// 	this.Redirect("/training", 302)
	// }
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

	this.Layout = "main_newemployee/layout/default.html"
	this.TplNames = "main_newemployee/article/newemployee_by_catalog.html"
}
