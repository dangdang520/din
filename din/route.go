package din

import (
	"fmt"
	"net/http"
)

type Router struct {
	HandlerFuncs map[string]HandlerFunc
}
func newRouter()*Router{
	return &Router{map[string]HandlerFunc{}}
}
func (r *Router) AddRoute(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	r.HandlerFuncs[key] = handlerFunc
}
func (r Router) Handle(resp http.ResponseWriter, req *http.Request)  {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := r.HandlerFuncs[key]; ok {
		handler(resp, req)
	} else {
		fmt.Fprintf(resp, "404 not found:%s \n", req.URL)
	}
}