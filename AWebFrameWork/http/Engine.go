package webServe

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(*Context)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// New is the constructor of Engine

func NewEngine() *Engine {
	return &Engine{
		router:  newRouter(),
	}
}


// GET method define
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.router.addRoute("GET", pattern, handler)
}

// POST method define
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.router.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router.handlers[key]; ok {
		handler( &Context{
			Writer: w,
			Req: req,
		})
	} else {
		fmt.Fprintf(w, "404 NOT FOUND %s\n", req.URL)
	}
}
