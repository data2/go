package main

import (
	"fmt"
	"html/template"
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

func name(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/index.html")
	data := map[string]string{
		"name": "data2",
	}
	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", webfun)
	http.HandleFunc("/name", name)
	fmt.Println("服务启动")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器开启错误：", err)
	}

}
