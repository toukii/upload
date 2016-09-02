package test

import (
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beego/httplib"
	"testing"
)

func TestT(t *testing.T) {

	b, err := ioutil.ReadFile("films.html")
	// b, err := httplib.Get("http://bookmark.daoapp.io").Bytes()
	if checkerr(err) {
		return
	}
	// req := httplib.Put("http://localhost:8080/upload/tests/upload_test.go")
	req := httplib.Put("http://upload.daoapp.io/upload/films/films.html")
	// req := httplib.Put("http://upload.daoapp.io/upload/第二题/Second.java")
	// req := httplib.Put("http://upload.daoapp.io/upload/小易打怪/Third.java")
	req.Body(b)
	resp, err := req.Response()
	if checkerr(err) {
		return
	}
	b2, err := ioutil.ReadAll(resp.Body)
	if checkerr(err) {
		return
	}
	fmt.Println(string(b2))
}

func checkerr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return true
	}
	return false
}

 curl -H 'Content-Type:text/plain' --data-binary @jsonTest.md "http://upload.daoapp.io/upload/jsonTest.md"