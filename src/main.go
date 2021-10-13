package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	// 创建一个Foo路由和处理函数
	//http.Handle("/foo", foohandler)
	println("start...")
	// 创建一个bar路由和处理函数
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})
	// 监听8080端口
	println("listening")
	log.Fatal(http.ListenAndServe(":8080", nil))
	println("start listening")
}
