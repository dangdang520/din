package din

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) AddRoute(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	e.router[key] = handlerFunc
}

func (e *Engine) Get(pattern string, handlerFunc HandlerFunc) {
	e.AddRoute("GET", pattern, handlerFunc)
}
func (e *Engine) Post(pattern string, handlerFunc HandlerFunc) {
	e.AddRoute("POST", pattern, handlerFunc)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Fprintf(w, "404 not found:%s \n", r.URL)
	}
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
