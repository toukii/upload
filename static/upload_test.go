package test

import (
	"fmt"
	"io/ioutil"

	"github.com/astaxie/beego/httplib"
	"testing"
)

func TestT(t *testing.T) {

	b, err := ioutil.ReadFile("Third.java")
	if checkerr(err) {
		return
	}
	req := httplib.Put("http://localhost:8080/upload/tests/upload_test.go")
	// req := httplib.Put("http://upload.daoapp.io/upload/第一题/First.java")
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
