package controllers

import (
	"fmt"
	"github.com/ulricqin/goutils/filetool"
	"star/g"
	"star/models"
	"star/models/blog"
	"star/models/catalog"
	"time"
)

//************************************************************************
const (
	NCATALOG_IMG_DIR = "static/uploads/catalogs"
)

type NCatalogController struct {
	NAdminController
}

func (this *NCatalogController) Add() {
	this.Data["IsAddCatalog"] = true
	this.Layout = "layout/admin.html"
	this.TplNames = "catalog/add.html"
}

func (this *NCatalogController) Edit() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("param id should be digit")
		return
	}

	c := catalog.NOneById(int64(id))
	if c == nil {
		this.Ctx.WriteString(fmt.Sprintf("no such catalog_id:%d", id))
		return
	}

	this.Data["Catalog"] = c
	this.Layout = "layout/admin.html"
	this.TplNames = "catalog/edit.html"
}

func (this *NCatalogController) Del() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("param id should be digit")
		return
	}

	c := catalog.NOneById(int64(id))
	if c == nil {
		this.Ctx.WriteString(fmt.Sprintf("no such catalog_id:%d", id))
		return
	}

	err = catalog.NDel(c)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.Ctx.WriteString("del success")
	return
}

func (this *NCatalogController) extractCatalog(imgMust bool) (*models.NewemployeeCatalog, error) {
	o := &models.NewemployeeCatalog{}
	o.Name = this.GetString("name")
	o.Ident = this.GetString("ident")
	o.Resume = this.GetString("resume")
	o.DisplayOrder = this.GetIntWithDefault("display_order", 0)

	if o.Name == "" {
		return nil, fmt.Errorf("name is blank")
	}

	if o.Ident == "" {
		return nil, fmt.Errorf("ident is blank")
	}

	_, header, err := this.GetFile("img")
	if err != nil && imgMust {
		return nil, err
	}

	if err == nil {
		ext := filetool.Ext(header.Filename)
		imgPath := fmt.Sprintf("%s/%s_%d%s", NCATALOG_IMG_DIR, o.Ident, time.Now().Unix(), ext)

		filetool.InsureDir(NCATALOG_IMG_DIR)
		err = this.SaveToFile("img", imgPath)
		if err != nil && imgMust {
			return nil, err
		}

		if err == nil {
			o.ImgUrl = "/" + imgPath
			if g.UseQiniu {
				if addr, er := g.UploadFile(imgPath, imgPath); er != nil {
					if imgMust {
						return nil, er
					}
				} else {
					o.ImgUrl = addr
					filetool.Unlink(imgPath)
				}
			}
		}
	}

	return o, nil
}

func (this *NCatalogController) DoEdit() {
	cid, err := this.GetInt("catalog_id")
	if err != nil {
		this.Ctx.WriteString("catalog_id is illegal")
		return
	}

	old := catalog.NOneById(int64(cid))
	if old == nil {
		this.Ctx.WriteString(fmt.Sprintf("no such catalog_id: %d", cid))
		return
	}

	var o *models.NewemployeeCatalog
	o, err = this.extractCatalog(false)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	old.Ident = o.Ident
	old.Name = o.Name
	old.Resume = o.Resume
	old.DisplayOrder = o.DisplayOrder
	if o.ImgUrl != "" {
		old.ImgUrl = o.ImgUrl
	}

	if err = catalog.NUpdate(old); err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.Redirect("/newemployee/blog", 302)

}

func (this *NCatalogController) DoAdd() {
	o, err := this.extractCatalog(true)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	_, err = catalog.NSave(o)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.Redirect("/newemployee/blog", 302)
}

//************************************************************************
type NAdminController struct {
	BaseController
}

func (this *NAdminController) CheckLogin() {
	if !this.IsAdmin {
		this.Redirect("/login", 302)
	}
}
func (this *NAdminController) Default() {
	this.Layout = "layout/admin.html"
	this.TplNames = "me/default.html"
}

//************************************************************************
type NArticleController struct {
	NAdminController
}

func (this *NArticleController) Draft() {
	var blogs []*models.NewemployeeBlog
	blog.NBlogs().Filter("Status", 0).All(&blogs)
	this.Data["Blogs"] = blogs
	this.Layout = "layout/admin.html"
	this.TplNames = "article/draft.html"
}

func (this *NArticleController) Add() {
	this.Data["Catalogs"] = catalog.NAll()
	this.Data["IsPost"] = true
	this.Layout = "layout/admin.html"
	this.TplNames = "article/add.html"
	this.JsStorage("deleteKey", "post/new")
}

func (this *NArticleController) DoAdd() {
	title := this.GetString("title")
	ident := this.GetString("ident")
	keywords := this.GetString("keywords")
	catalog_id := this.GetIntWithDefault("catalog_id", -1)
	aType := this.GetIntWithDefault("type", -1)
	status := this.GetIntWithDefault("status", -1)
	content := this.GetString("content")

	if catalog_id == -1 || aType == -1 || status == -1 {
		this.Ctx.WriteString("catalog || type || status is illegal")
		return
	}

	if title == "" || ident == "" {
		this.Ctx.WriteString("title or ident is blank")
		return
	}

	cp := catalog.NOneById(int64(catalog_id))
	if cp == nil {
		this.Ctx.WriteString("catalog_id not exists")
		return
	}

	b := &models.NewemployeeBlog{Ident: ident, Title: title, Keywords: keywords, CatalogId: int64(catalog_id), Type: int8(aType), Status: int8(status)}
	_, err := blog.NSave(b, content)

	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.JsStorage("deleteKey", "post/new")
	this.Redirect("/newemployee/catalog/"+cp.Ident, 302)

}

func (this *NArticleController) Edit() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("get param id fail")
		return
	}

	b := blog.NOneById(int64(id))
	if b == nil {
		this.Ctx.WriteString("no such article")
		return
	}

	this.Data["Content"] = blog.NReadBlogContent(b).Content
	this.Data["Blog"] = b
	this.Data["Catalogs"] = catalog.NAll()
	this.Layout = "layout/admin.html"
	this.TplNames = "article/edit.html"
}

func (this *NArticleController) DoEdit() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("get param id fail")
		return
	}

	b := blog.NOneById(int64(id))
	if b == nil {
		this.Ctx.WriteString("no such article")
		return
	}

	title := this.GetString("title")
	ident := this.GetString("ident")
	keywords := this.GetString("keywords")
	catalog_id := this.GetIntWithDefault("catalog_id", -1)
	aType := this.GetIntWithDefault("type", -1)
	status := this.GetIntWithDefault("status", -1)
	content := this.GetString("content")

	if catalog_id == -1 || aType == -1 || status == -1 {
		this.Ctx.WriteString("catalog || type || status is illegal")
		return
	}

	if title == "" || ident == "" {
		this.Ctx.WriteString("title or ident is blank")
		return
	}

	cp := catalog.NOneById(int64(catalog_id))
	if cp == nil {
		this.Ctx.WriteString("catalog_id not exists")
		return
	}

	b.Ident = ident
	b.Title = title
	b.Keywords = keywords
	b.CatalogId = int64(catalog_id)
	b.Type = int8(aType)
	b.Status = int8(status)

	err = blog.NUpdate(b, content)

	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.JsStorage("deleteKey", "post/edit")
	this.Redirect("/newemployee/catalog/"+cp.Ident, 302)
}

func (this *NArticleController) Del() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("get param id fail")
		return
	}

	b := blog.NOneById(int64(id))
	if b == nil {
		this.Ctx.WriteString("no such article")
		return
	}

	err = blog.NDel(b)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.Ctx.WriteString("del success")
	return
}

//************************************************************************
