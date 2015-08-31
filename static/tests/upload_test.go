package test

import (
	"fmt"
	"io/ioutil"

	"testing"
	"github.com/astaxie/beego/httplib"
)

func TestT(t *testing.T) {

	b, err := ioutil.ReadFile("main.go")
	if checkerr(err) {
		return
	}
	req := httplib.Put("http://localhost:8080/upload/123")
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
