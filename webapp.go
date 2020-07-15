package main

import (
	"fmt"
	"net/http"
)

func webfun(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, "开启go web")
	for k, v := range r.URL.Query() {
		fmt.Println("key:", k, ", value:", v[0])
	}
	for k, v := range r.PostForm {
		fmt.Println("key:", k, ", value:", v[0])
	}

}

func main() {
	http.HandleFunc("/", webfun)
	fmt.Println("服务启动")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器开启错误：", err)
	}

}
