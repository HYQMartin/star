package controllers

import (
	"fmt"
	"io"
	"os"
	"path"
	"star/models"
	"star/models/blog"
	"star/models/catalog"
	"strings"

	"code.google.com/p/go-uuid/uuid"
)

//************************************************************************
const (
	CATALOG_IMG_DIR = "static/upload/catalogs"
)

type CatalogController struct {
	AdminController
}

func (this *CatalogController) Add() {
	if !this.IsAdmin {
		this.Ctx.WriteString("you are not the admin")
		return
	}
	var p string
	cident := this.GetString("cident")
	for i := 0; i < len(FIXPAGE); i++ {
		if cident == FIXPAGE[i] {
			p = "/"
		}
	}
	if p == "" {
		p = "/catalog/" + cident
	}
	this.Data["ReturnPath"] = p
	this.Data["IsAddCatalog"] = true
	this.Layout = "layout/admin.html"
	this.TplNames = "catalog/add.html"
}

func (this *CatalogController) Edit() {
	if !this.IsAdmin {
		this.Ctx.WriteString("you are not the admin")
		return
	}

	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("param id should be digit")
		return
	}

	c := catalog.OneById(int64(id))
	if c == nil {
		this.Ctx.WriteString(fmt.Sprintf("no such catalog_id:%d", id))
		return
	}

	var p string
	cident := this.GetString("cident")
	for i := 0; i < len(FIXPAGE); i++ {
		if cident == FIXPAGE[i] {
			p = "/"
		}
	}
	if p == "" {
		p = "/catalog/" + cident
	}
	this.Data["ReturnPath"] = p
	this.Data["Catalog"] = c
	this.Layout = "layout/admin.html"
	this.TplNames = "catalog/edit.html"
}

func (this *CatalogController) Del() {
	if !this.IsAdmin {
		this.Ctx.WriteString("you are not the admin")
		return
	}
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("param id should be digit")
		return
	}

	c := catalog.OneById(int64(id))
	if c == nil {
		this.Ctx.WriteString(fmt.Sprintf("no such catalog_id:%d", id))
		return
	}

	err = catalog.Del(c)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.Ctx.WriteString("del success")
	return
}

func (this *CatalogController) extractCatalog(imgMust bool) (*models.Catalog, error) {
	o := &models.Catalog{}
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

	file, header, err := this.GetFile("img")
	if err != nil && imgMust {
		return nil, err
	}
	defer file.Close()

	if err == nil {
		filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)
		err = os.MkdirAll(path.Join("static", "upload"), 0775)
		if err != nil {
			this.Ctx.WriteString(err.Error())
			return nil, err
		}
		outFile, err := os.Create(path.Join("static", "upload", filename))
		if err != nil {
			this.Ctx.WriteString(err.Error())
			return nil, err
		}
		defer outFile.Close()
		io.Copy(outFile, file)

		imgPath := fmt.Sprintf("/static/upload/%s", filename)
		o.ImgUrl = imgPath
	}

	return o, nil
}

func (this *CatalogController) DoEdit() {
	if !this.IsAdmin {
		this.Ctx.WriteString("you are not the admin")
		return
	}
	cid, err := this.GetInt("catalog_id")
	if err != nil {
		this.Ctx.WriteString("catalog_id is illegal")
		return
	}

	old := catalog.OneById(int64(cid))
	if old == nil {
		this.Ctx.WriteString(fmt.Sprintf("no such catalog_id: %d", cid))
		return
	}

	var o *models.Catalog
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

	if err = catalog.Update(old); err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.Redirect("/", 302)

}

func (this *CatalogController) DoAdd() {
	if !this.IsAdmin {
		this.Ctx.WriteString("you are not the admin")
		return
	}
	o, err := this.extractCatalog(true)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	_, err = catalog.Save(o)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.Redirect("/", 302)
}

//************************************************************************
type AdminController struct {
	BaseController
}

func (this *AdminController) CheckLogin() {
	if !this.IsAdmin && !this.IsLogin {
		this.Redirect("/login", 302)
	}
}
func (this *AdminController) Default() {
	this.Layout = "layout/admin.html"
	this.TplNames = "me/default.html"
}

//************************************************************************
type ArticleController struct {
	AdminController
}

func (this *ArticleController) Draft() {
	var blogs []*models.Blog
	blog.Blogs().Filter("Status", 0).All(&blogs)
	this.Data["ReturnPath"] = "/catalog/share"
	this.Data["Blogs"] = blogs
	this.Layout = "layout/admin.html"
	this.TplNames = "article/draft.html"
}

func (this *ArticleController) Add() {
	tmp := catalog.All()
	t := make([]*models.Catalog, 0)
	flag := true
	for i := 0; i < len(tmp); i++ {
		for j := 0; j < len(FIXPAGE); j++ {
			if tmp[i].Ident == FIXPAGE[j] {
				flag = false
				break
			}
		}
		if flag {
			t = append(t, tmp[i])
		}
		flag = true
	}
	this.Data["Catalogs"] = t
	this.Data["IsPost"] = true

	var p string
	cident := this.GetString("cident")
	for i := 0; i < len(FIXPAGE); i++ {
		if cident == FIXPAGE[i] {
			p = "/"
		}
	}
	if p == "" {
		p = "/catalog/" + cident
	}
	this.Data["ReturnPath"] = p
	this.Layout = "layout/admin.html"
	this.TplNames = "article/add.html"
	this.JsStorage("deleteKey", "post/new")
}

func (this *ArticleController) DoAdd() {
	if !this.IsLogin && !this.IsAdmin {
		this.Ctx.WriteString("you are not the admin or author")
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

	cp := catalog.OneById(int64(catalog_id))
	if cp == nil {
		this.Ctx.WriteString("catalog_id not exists")
		return
	}

	b := &models.Blog{Ident: ident, Title: title, Keywords: keywords, CatalogId: int64(catalog_id), Type: int8(aType), Status: int8(status), Creator: this.UserName}

	_, err := blog.Save(b, content)

	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.JsStorage("deleteKey", "post/new")
	this.Redirect("/catalog/"+cp.Ident, 302)
	this.Ctx.WriteString("success")
	return

}

func (this *ArticleController) Edit() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("get param id fail")
		return
	}

	b := blog.OneById(int64(id))
	if b == nil {
		this.Ctx.WriteString("no such article")
		return
	}
	if this.UserName != b.Creator && !this.IsAdmin {
		this.Ctx.WriteString("you are not author or admin")
		return
	}
	var p string
	cident := this.GetString("cident")
	for i := 0; i < len(FIXPAGE); i++ {
		if cident == FIXPAGE[i] {
			p = "/"
		}
	}
	if p == "" {
		p = "/catalog/" + cident
	}
	this.Data["ReturnPath"] = p

	this.Data["Content"] = blog.ReadBlogContent(b).Content
	this.Data["Blog"] = b

	tmp := catalog.All()
	t := make([]*models.Catalog, 0)
	flag := true
	for i := 0; i < len(tmp); i++ {
		for j := 0; j < len(FIXPAGE); j++ {
			if tmp[i].Ident == FIXPAGE[j] {
				flag = false
				break
			}
		}
		if flag {
			t = append(t, tmp[i])
		}
		flag = true
	}
	this.Data["Catalogs"] = t
	this.Layout = "layout/admin.html"
	this.TplNames = "article/edit.html"
}

func (this *ArticleController) DoEdit() {
	if !this.IsLogin && !this.IsAdmin {
		this.Ctx.WriteString("you are not the admin or author")
		return
	}
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("get param id fail")
		return
	}

	b := blog.OneById(int64(id))
	if b == nil {
		this.Ctx.WriteString("no such article")
		return
	}
	if this.UserName != b.Creator && !this.IsAdmin {
		this.Ctx.WriteString("you are not author or admin")
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

	cp := catalog.OneById(int64(catalog_id))
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

	err = blog.Update(b, content)

	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	this.Ctx.WriteString("success")
	this.JsStorage("deleteKey", "post/edit")
	this.Redirect("/catalog/"+cp.Ident, 302)
	return
}

func (this *ArticleController) Del() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.WriteString("get param id fail")
		return
	}

	b := blog.OneById(int64(id))
	if b == nil {
		this.Ctx.WriteString("no such article")
		return
	}
	if this.UserName != b.Creator && !this.IsAdmin {
		this.Ctx.WriteString("you are not author or admin")
		return
	}
	err = blog.Del(b)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	cident := this.GetString("cident")
	this.Redirect("/catalog/"+cident, 302)
	// this.Ctx.WriteString("del success")
	return
}

//************************************************************************
