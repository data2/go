package main

import (
	"fmt"
	"net/http"
)

func webfun(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "开启go web")
}

func main() {
	http.HandleFunc("/", webfun)
	fmt.Println("服务启动")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("服务器开启错误：", err)
	}

}
