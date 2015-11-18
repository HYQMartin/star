package controllers

import (
	"github.com/astaxie/beego"
	/*"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"*/)

type FileController struct {
	beego.Controller
}

func (c *FileController) Get() {
	/*if context.Method == "DELETE" {
		id := context.Int("id")
		model.RemoveFile(id)
		Json(context, true).End()
		context.Do("attach_delete", id)
		return
	}*/
	c.TplNames = "File.html"
	var files, pager map[string]string
	files = make(map[string]string)
	pager = make(map[string]string)
	files["file1"] = "file1"
	files["file2"] = "file2"
	pager["pager1"] = "pager1"

	c.Data["Title"] = "媒体文件"
	c.Data["Files"] = files
	c.Data["Pager"] = pager
	c.Data["Name"] = "Name"

	return
}

type FileUploadController struct {
	beego.Controller
}

func (c *FileUploadController) Get() {
	/*
		var req *http.Request
		req = context.Request
		req.ParseMultipartForm(32 << 20)
		f, h, e := req.FormFile("file")
		if e != nil {
			Json(context, false).Set("msg", e.Error()).End()
			return
		}
		data, _ := ioutil.ReadAll(f)
		maxSize := context.App().Config().Int("app.upload_size")
		defer func() {
			f.Close()
			data = nil
			h = nil
		}()
		if len(data) >= maxSize {
			Json(context, false).Set("msg", "文件应小于10M").End()
			return
		}
		if !strings.Contains(context.App().Config().String("app.upload_files"), path.Ext(h.Filename)) {
			Json(context, false).Set("msg", "文件只支持Office文件，图片和zip存档").End()
			return
		}
		ff := new(model.File)
		ff.Name = h.Filename
		ff.Type = context.StringOr("type", "image")
		ff.Size = int64(len(data))
		ff.ContentType = h.Header["Content-Type"][0]
		ff.Author, _ = strconv.Atoi(context.Cookie("token-user"))
		ff.Url = model.CreateFilePath(context.App().Get("upload_dir"), ff)
		e = ioutil.WriteFile(ff.Url, data, os.ModePerm)
		if e != nil {
			Json(context, false).Set("msg", e.Error()).End()
			return
		}
		model.CreateFile(ff)
		Json(context, true).Set("file", ff).End()
		context.Do("attach_created", ff)
	*/
	return
}
