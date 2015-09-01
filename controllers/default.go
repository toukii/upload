package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// @router / [get]
func (c *MainController) LoadUpload() {
	c.TplNames = "upload.html"
}

// @router /uploadform [post]
func (c *MainController) UploadForm() {
	_, file, err := c.GetFile("filename")
	if nil == err {
		if serr := c.SaveToFile("filename", "./static/"+file.Filename); serr == nil {
		} else {
			beego.Error(serr)
			c.Ctx.WriteString(serr.Error())
		}
		c.Ctx.ResponseWriter.Write([]byte("http://localhost:8080/download/" + file.Filename))
		return
	}
	beego.Error(err)
	c.Ctx.WriteString(err.Error())
}

// @router /download/* [get]
func (c *MainController) Download() {
	filename := c.Ctx.Input.Param(":splat")
	beego.Debug(filename)
	dstfilename := "./static/" + filename
	c.Ctx.Output.Download(dstfilename, filename)
}

// @router /loadfile/* [get]
func (c *MainController) LoadFile() {
	filename := c.Ctx.Input.Param(":splat")
	beego.Debug(filename)
	if file, err := os.Open("./static/" + filename); err != nil {
		beego.Error(err)
		c.Ctx.WriteString(err.Error())
	} else {
		if b, err := ioutil.ReadAll(file); err != nil {
			beego.Error(err)
			c.Ctx.WriteString(err.Error())
		} else {
			c.Ctx.Output.Body(b)
		}
	}
}

// @router /list/* [get]
func (c *MainController) ListFile() {
	beego.Debug(c.Ctx.Request.RequestURI)
	pathname := c.Ctx.Input.Param(":splat")
	beego.Debug(pathname)
	if "a" == pathname {
		pathname = ""
	}
	fs, err := ioutil.ReadDir("./static/" + pathname)
	if checkerr(err) {
		c.Ctx.WriteString(err.Error())
	}
	dirs := make([]string, 0, len(fs))
	files := make([]string, 0, len(fs))
	for _, it := range fs {
		if it.IsDir() {
			dirs = append(dirs, filepath.Join(pathname, it.Name()))
			continue
		}
		files = append(files, filepath.Join(pathname, it.Name()))
	}
	c.Data["dirs"] = dirs
	c.Data["files"] = files
	c.TplNames = "list.html"
}

// @router /delfile/* [*]
func (c *MainController) DeleteFile() {
	file := c.Ctx.Input.Param(":splat")
	beego.Debug(file)
	err := os.RemoveAll("./static/" + file)
	if checkerr(err) {
		c.Ctx.WriteString(err.Error())
	}
	c.Ctx.WriteString(`{"ret":"success"}`)
}

// @router /upload/* [*]
func (c *MainController) Upload() {
	rw := c.Ctx.ResponseWriter
	req := c.Ctx.Request
	if req.Method == "GET" {
		rw.Write([]byte(""))
	}
	req.ParseForm()
	length := req.Header.Get("Content-Length")
	fmt.Println(length)
	b, err := ioutil.ReadAll(req.Body)
	if checkerr(err) {
		rw.Write([]byte("error"))
	}
	filename := c.Ctx.Input.Param(":splat")
	beego.Debug(filename)
	file, err := os.OpenFile("./static/"+filename, os.O_CREATE|os.O_WRONLY, 0644)
	if checkerr(err) {
		rw.Write([]byte("error"))
	}
	_, err = file.Write(b)
	if checkerr(err) {
		rw.Write([]byte("error"))
	}
	rw.Write([]byte("http://localhost:8080/download/" + filename))
}

func checkerr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
