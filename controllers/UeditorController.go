package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "html/template"
	"io"
	"log"
	// "net/http"
	"os"
	"path"
	"strings"

	"code.google.com/p/go-uuid/uuid"
)

var configJson []byte // 当客户端请求 /ueditor/go/controller?action=config 返回的json内容

func init() {
	file, err := os.Open("conf/config.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer file.Close()
	buf := bytes.NewBuffer(nil)
	buf.ReadFrom(file)

	configJson = buf.Bytes()
}

type UeditorController struct {
	BaseController
}

func (this *UeditorController) Get() {
	if !this.IsAdmin && !this.IsLogin {
		this.Redirect("/login", 302)
	}

	action := this.GetString("action")
	if action == "config" {
		this.Ctx.WriteString(string(configJson))
	}
}
func (this *UeditorController) Post() {
	if !this.IsAdmin && !this.IsLogin {
		this.Redirect("/login", 302)
	}

	action := this.GetString("action")
	if action != "uploadimage" {
		this.Ctx.WriteString("action != uploadimage")
		return
	}

	file, header, err := this.GetFile("upfile")
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	defer file.Close()

	filename := strings.Replace(uuid.NewUUID().String(), "-", "", -1) + path.Ext(header.Filename)

	err = os.MkdirAll(path.Join("static", "upload"), 0775)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	outFile, err := os.Create(path.Join("static", "upload", filename))
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	defer outFile.Close()

	io.Copy(outFile, file)

	b, err := json.Marshal(map[string]string{
		"url":      fmt.Sprintf("/static/upload/%s", filename), //保存后的文件路径
		"title":    "",                                         //文件描述，对图片来说在前端会添加到title属性上
		"original": header.Filename,                            //原始文件名
		"state":    "SUCCESS",                                  //上传状态，成功时返回SUCCESS,其他任何值将原样返回至图片上传框中
	})
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	this.Ctx.WriteString(string(b))
	return
}
