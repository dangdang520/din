package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func FooHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

type Engine struct {
}

func (e *Engine) ServeHTTP(rep http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/foo":
		fmt.Fprintf(rep, "%q", req.URL.Path)
		break
	case "/test":
		fmt.Fprintf(rep, "%q", "444555")
		break
	default:
		fmt.Fprintf(rep, "%q", "默认数据")
	}
}

func main() {
	//http.HandleFunc("/foo", FooHandle)
	eng := new(Engine)
	log.Fatal(http.ListenAndServe(":8080",
		eng))
}
