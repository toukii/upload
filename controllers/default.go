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

var (
	// volumn = "/usr/static/"
	volumn = "./static/"
)

type MainController struct {
	beego.Controller
}

// @router / [get]
func (c *MainController) Home() {
	c.Redirect("/list/_home", 302)
}

// @router /list/ [get]
func (c *MainController) LHome() {
	c.Redirect("/list/_home", 302)
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
		if serr := c.SaveToFile("filename", volumn+file.Filename); serr == nil {
		} else {
			beego.Error(serr)
			c.Ctx.WriteString(serr.Error())
		}
		c.Redirect("/list/_home", 302)
	}
	beego.Error(err)
	c.Redirect("/list/_home", 302)
}

// @router /uploadform/* [post]
func (c *MainController) DirUploadForm() {
	dir := c.Ctx.Input.Param(":splat")
	_, file, err := c.GetFile("filename")
	if nil == err {
		if serr := c.SaveToFile("filename", volumn+dir+"/"+file.Filename); serr == nil {
		} else {
			beego.Error(serr)
			c.Ctx.WriteString(serr.Error())
		}
		c.Redirect("/list/"+dir, 302)
	}
	beego.Error(err)
	c.Redirect("/list/"+dir, 302)
}

// @router /download/* [get]
func (c *MainController) Download() {
	filename := c.Ctx.Input.Param(":splat")
	beego.Debug(filename)
	dstfilename := volumn + filename
	c.Ctx.Output.Download(dstfilename, filename)
}

// @router /loadfile/* [get]
func (c *MainController) LoadFile() {
	filename := c.Ctx.Input.Param(":splat")
	beego.Debug(filename)
	if file, err := os.Open(volumn + filename); err != nil {
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

// @router /display/* [post]
func (c *MainController) PostDisplay() {
	file := c.Ctx.Input.Param(":splat")
	req := c.Ctx.Request
	req.ParseForm()
	content := req.Form.Get("content")
	createFile(file, content)
	c.Redirect("/display/"+file, 302)
}

// @router /display/* [get]
func (c *MainController) Display() {
	filename := c.Ctx.Input.Param(":splat")
	fileview := FileView{Name: filename}
	filetypes := strings.Split(filename, ".")
	imged := false
	if len(filetypes) > 1 {
		if strings.Contains(".pdf.doc.docx", filetypes[len(filetypes)-1]) {
			c.Redirect("/loadfile/"+filename, 302)
		}
		if strings.Contains(imgs, filetypes[len(filetypes)-1]) {
			fileview.Img = "/loadfile/" + filename
			imged = true
		}
	}
	if !imged {
		info, err := os.Stat(volumn + filename)
		if nil == err && info.Size() < 1e6 {
			fileview.Content = readFile(filename)
			if len(fileview.Content) < 1 {
				fileview.Content = " "
			}
		}

	}
	c.Data["dir"] = filepath.Dir(filename)
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
	if "_home" == pathname {
		pathname = ""
	}
	fs, err := ioutil.ReadDir(volumn + pathname)
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
			if strings.Contains(imgs, filetypes[len(filetypes)-1]) {
				fileview.Img = "/loadfile/" + name
			}
		}
		fileviews = append(fileviews, fileview)
	}
	c.Data["dir"] = pathname
	c.Data["dirs"] = dirs
	c.Data["fileviews"] = fileviews
	c.TplNames = "list.html"
}

// @router /delfile/* [*]
func (c *MainController) DeleteFile() {
	beego.Info(c.Ctx.Request.RemoteAddr)
	file := c.Ctx.Input.Param(":splat")
	beego.Debug(file)
	now := goutils.LocNow("Asia/Shanghai")
	if now.Second()%10 < 3 {
		err := os.RemoveAll(volumn + file)
		if checkerr(err) {
			c.Ctx.WriteString(file)
			return
		}
	}
	dir := filepath.Dir(file)
	if len(dir) <= 1 {
		dir = "_home"
	}
	c.Ctx.WriteString(dir)
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

// @router /topic/* [get]
func (c *MainController) GTopics() {
	c.Data["dir"] = c.Ctx.Input.Param(":splat")
	c.TplNames = "topic.html"
}

// @router /topic/* [post]
func (c *MainController) PTopics() {
	dir := c.Ctx.Input.Param(":splat")
	req := c.Ctx.Request
	req.ParseForm()
	title := req.Form.Get("title")
	content := req.Form.Get("content")
	createFile(dir+"/"+title, content)
	fmt.Println(dir, filepath.Dir(dir))
	c.Redirect("/list/"+dir, 302)
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
	_, err := os.Stat(volumn + dir)
	if !checkerr(err) {
		os.Remove(volumn + filename)
	}
	os.MkdirAll(volumn+dir, 0777)
	file, err := os.OpenFile(volumn+filename, os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	if checkerr(err) {
		return err
	}
	_, err = file.WriteString(content)
	return err
}

func readFile(filename string) string {
	file, err := os.Open(volumn + filename)
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
