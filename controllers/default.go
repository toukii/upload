package controllers

import (
	"fmt"
	"github.com/shaalx/goutils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// @router / [get]
func (c *MainController) Home() {
	// c.TplNames = "upload.html"
	c.Redirect("/list/a", 302)
}

// @router /upload [get]
func (c *MainController) LoadUpload() {
	c.TplNames = "upload.html"
}

// @router /upload/* [get]
func (c *MainController) LoadUploads() {
	dir := c.Ctx.Input.Param(":splat")
	c.Data["dir"] = dir
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
		c.Redirect("/list/a", 302)
	}
	beego.Error(err)
	c.Ctx.WriteString(err.Error())
}

// @router /uploadform/* [post]
func (c *MainController) DirUploadForm() {
	dir := c.Ctx.Input.Param(":splat")
	_, file, err := c.GetFile("filename")
	if nil == err {
		if serr := c.SaveToFile("filename", "./static/"+dir+"/"+file.Filename); serr == nil {
		} else {
			beego.Error(serr)
			c.Ctx.WriteString(serr.Error())
		}
		c.Redirect("/list/"+dir, 302)
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
		defer file.Close()
		if b, err := ioutil.ReadAll(file); err != nil {
			beego.Error(err)
			c.Ctx.WriteString(err.Error())
		} else {
			c.Ctx.Output.Body(b)
		}
	}
}

// @router /display/* [get]
func (c *MainController) Display() {
	filename := c.Ctx.Input.Param(":splat")
	fileview := FileView{Name: filename}
	filetypes := strings.Split(filename, ".")
	fmt.Println(filetypes)
	imged := false
	if len(filetypes) > 1 {
		if strings.Contains(".pdf", filetypes[1]) {
			c.Redirect("/loadfile/"+filename, 302)
		}
		if strings.Contains(imgs, filetypes[1]) {
			fileview.Img = "/loadfile/" + filename
			imged = true
		}
	}
	if !imged {
		info, err := os.Stat("./static/" + filename)
		fmt.Println(info, filename)
		if nil == err && info.Size() < 1e6 {
			fileview.Content = readFile(filename)
		}
	}
	fmt.Println(fileview)
	c.Data["file"] = fileview
	c.TplNames = "display.html"
}

type FileView struct {
	Name    string
	Content string
	Img     string
}

var (
	imgs = "png,gif,jpg,jpeg,bmp,tiff"
)

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
	fileviews := make([]FileView, 0, len(fs))
	for _, it := range fs {
		if it.IsDir() {
			dirs = append(dirs, filepath.Join(pathname, it.Name()))
			continue
		}
		name := filepath.Join(pathname, it.Name())
		fileview := FileView{Name: name}
		filetypes := strings.Split(name, ".")
		if len(filetypes) > 1 {
			if strings.Contains(imgs, filetypes[1]) {
				fileview.Img = "/loadfile/" + name
			}
		}
		fileviews = append(fileviews, fileview)
	}
	c.Data["dirs"] = dirs
	c.Data["fileviews"] = fileviews
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

// @router /upload/* [post,put]
func (c *MainController) Upload() {
	rw := c.Ctx.ResponseWriter
	req := c.Ctx.Request
	// if req.Method == "GET" {
	// 	return
	// }
	b, err := ioutil.ReadAll(req.Body)
	if checkerr(err) {
		rw.Write([]byte(err.Error()))
	}
	filename := c.Ctx.Input.Param(":splat")
	createFile(filename, goutils.ToString(b))
}

// @router /topic [get]
func (c *MainController) GTopic() {
	c.TplNames = "topic.html"
}

// @router /topic [post]
func (c *MainController) PTopic() {
	req := c.Ctx.Request
	req.ParseForm()
	title := req.Form.Get("title")
	content := req.Form.Get("content")
	createFile(title, content)
	c.Redirect("/", 302)
}

func createFile(filename, content string) error {
	dir := filepath.Dir(filename)
	_, err := os.Stat("./static/" + dir)
	if checkerr(err) {
		os.MkdirAll("./static/"+dir, 0777)
	}
	file, err := os.OpenFile("./static/"+filename, os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if checkerr(err) {
		return err
	}
	_, err = file.WriteString(content)
	return err
}

func readFile(filename string) string {
	file, err := os.Open("./static/" + filename)
	defer file.Close()
	if err != nil {
		return err.Error()
	} else {
		b, err := ioutil.ReadAll(file)
		if err != nil {
			return err.Error()
		}
		return goutils.ToString(b)
	}
}

func checkerr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}
