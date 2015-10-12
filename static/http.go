package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"time"
)

func getOnce(rw http.ResponseWriter, req *http.Request) {
	resp, err := http.Get("http://greeting-shaalx.myalauda.cn/greeting")
	if checkErr(err) {
		return
	}
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))

	var content map[string]interface{}
	err = json.Unmarshal(b, &content)
	if checkErr(err) {
		return
	}

	t := template.New("index.html")
	t, err = t.ParseFiles("index.html")
	if checkErr(err) {
		return
	}
	t.Execute(rw, content)
}

func get() {
	for {
		resp, err := http.Get("http://greeting-shaalx.myalauda.cn/greeting")
		if checkErr(err) {
			return
		}
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
		fmt.Print(".")
	}
}
func main() {
	// getOnce()

	// for i := 0; i < 5; i++ {
	// 	go get()
	// }

	http.HandleFunc("/", getOnce)
	http.ListenAndServe(":8080", nil)

	time.Sleep(time.Second)

}

func checkErr(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}
