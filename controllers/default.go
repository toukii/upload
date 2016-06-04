package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/everfore/exc"
	"github.com/shaalx/goutils"
	// "html/template"
	"github.com/everfore/rpcsv"
	"io/ioutil"
	"net/http"
	"net/rpc"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	volumn = "/usr/static/upload/"
	// volumn = "./static/"
	excm = exc.NewCMD("ls")

	RPC_Client *rpc.Client
	// rpc_tcp_server = "127.0.0.1:8800"
	rpc_tcp_server = "upload.t0.daoapp.io:61441"
)

func connect() *rpc.Client {
	return rpcsv.RPCClientWithCodec(rpc_tcp_server)
}

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
	c.TplName = "upload.html"
}

// @router /upload/* [get]
func (c *MainController) LoadUploads() {
	dir := c.Ctx.Input.Param(":splat")
	c.Data["dir"] = dir
	c.TplName = "upload.html"
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
	filetype := path.Ext(filename)
	imged := false
	if strings.Contains(docs, filetype) {
		c.Redirect("/loadfile/"+filename, 302)
	}
	if strings.Contains(imgs, filetype) {
		fileview.Img = "/loadfile/" + filename
		imged = true
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
	line := readLine(volumn + filename)
	/*if strings.Count(line, "http://") == 1 || strings.Count(line, "https://") == 1 {
		// fileview.URI = template.HTML(goutils.ToString(LoadURL(line)))
		fileview.URI = line
	}*/
	fileview.Content = line
	c.Data["dir"] = filepath.Dir(filename)
	c.Data["file"] = fileview
	c.TplName = "display.html"
}

func readLine(filename string) string {
	file, err := os.Open(filename)
	defer file.Close()
	if checkerr(err) {
		return ""
	}
	b := make([]byte, 300)
	n, err := file.Read(b)
	if checkerr(err) {
		return ""
	}
	return goutils.ToString(b[:n])
}

func LoadURL(uri string) []byte {
	resp, err := http.Get(uri)
	if checkerr(err) {
		return nil
	}
	b, err := ioutil.ReadAll(resp.Body)
	if checkerr(err) {
		return nil
	}
	return b
}

type FileView struct {
	Name    string
	Content string
	Img     string
	URI     interface{}
}

var (
	imgs = ".png.gif.jpg.jpeg.bmp.tiff"
	docs = ".pdf.doc.docx.html"
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
		filetype := path.Ext(name)
		if strings.Contains(imgs, filetype) {
			fileview.Img = "/loadfile/" + name
		}
		fileviews = append(fileviews, fileview)
	}
	c.Data["dir"] = pathname
	c.Data["dirs"] = dirs
	c.Data["fileviews"] = fileviews
	c.TplName = "list.html"
}

// @router /delfile/* [*]
func (c *MainController) DeleteFile() {
	beego.Info(c.Ctx.Request.RemoteAddr)
	file := c.Ctx.Input.Param(":splat")
	beego.Debug(file, path.Ext(file))
	ext := path.Ext(file)
	if len(ext) > 0 && strings.Contains(imgs, path.Ext(file)) {
		return
	}
	/*inputName := c.GetString("Name")
	fmt.Println(file, inputName, file == inputName)
	if file != inputName {
		c.Ctx.WriteString("_home")
		return
	}*/
	/*now := goutils.LocNow("Asia/Shanghai")
	if now.Second()%10 < 3 {}
	*/
	err := os.RemoveAll(volumn + file)
	if checkerr(err) {
		c.Ctx.WriteString(file)
		return
	}

	dir := filepath.Dir(file)
	if "." == dir {
		dir = "/"
	}
	fmt.Println(dir)
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

// @router /job [get]
func (c *MainController) GJob() {
	fmt.Println(c.Ctx.Request.RequestURI)
	c.Data["name"] = c.Ctx.Request.Host
	c.TplName = "job.html"
}

// @router /job/* [get]
func (c *MainController) GJobs() {
	c.Data["dir"] = c.Ctx.Input.Param(":splat")
	fmt.Println(c.Ctx.Request.RequestURI)
	c.Data["name"] = c.Ctx.Request.Host
	c.TplName = "job.html"
}

// @router /job/* [post]
func (c *MainController) PJobs() string {
	// dir := c.Ctx.Input.Param(":splat")
	req := c.Ctx.Request
	req.ParseForm()
	title := req.Form.Get("name")
	content := req.Form.Get("target")
	fmt.Println(title, content)
	c.Data["result"] = fmt.Sprintf("%s,%s", title, content)
	return fmt.Sprintf("%s,%s", title, content)
}

// @router /job [post]
func (c *MainController) PJob() {
	req := c.Ctx.Request
	req.ParseForm()
	title := req.Form.Get("name")
	content := req.Form.Get("target")
	fmt.Printf("%s,%s", title, content)
	job := rpcsv.Job{Name: title, Target: content}
	RPC_Client = connect()
	b := make([]byte, 10)
	err := RPC_Client.Call("RPC.Job", &job, &b)
	if goutils.CheckErr(err) {
		c.Data["json"] = err
	} else {
		// fmt.Println(goutils.ToString(b))
		c.Data["json"] = goutils.ToString(b)
	}
	c.ServeJSON(true)
}

// @router /topic [get]
func (c *MainController) GTopic() {
	c.TplName = "topic.html"
}

// @router /topic/* [get]
func (c *MainController) GTopics() {
	c.Data["dir"] = c.Ctx.Input.Param(":splat")
	c.TplName = "topic.html"
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

// @router /bash [get]
func (c *MainController) Bash() {
	c.TplName = "bash.html"
}

// @router /bash [post]
func (c *MainController) PBash() {
	req := c.Ctx.Request
	req.ParseForm()
	shcont := req.Form.Get("shcont")
	// fmt.Println(c.Ctx.Input.Request)
	// fmt.Println(c.Ctx.Input.Data)
	// shcont := c.Ctx.Input.Param("shcont")
	beego.Info(shcont)
	excm.Reset(shcont)
	b, err := excm.Debug().Do()
	if checkerr(err) {
		// c.Data["err"] = err.Error()
		b = append(b, goutils.ToByte(err.Error())...)
	} else {
		// c.Data["result"] = strings.Split(goutils.ToString(b), " ")
	}
	c.Ctx.ResponseWriter.Write(b)
	// return fmt.Sprintf("%v", ret)
	// c.TplName = "bash.html"
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
