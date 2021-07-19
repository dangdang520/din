package din

import (
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) Get(pattern string, handlerFunc HandlerFunc) {
	e.router.AddRoute("GET", pattern, handlerFunc)
}
func (e *Engine) Post(pattern string, handlerFunc HandlerFunc) {
	e.router.AddRoute("POST", pattern, handlerFunc)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	e.router.Handle(w,r)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
