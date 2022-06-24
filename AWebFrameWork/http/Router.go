package webServe

import (
	"log"
	"net/http"
)

// Router module
type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

// add a router  make something fundamental for get/post
func (router *router) addRoute(method, pattern string, handlerFunc HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	router.handlers[key] = handlerFunc
}

func (router *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handlerFunc, ok := router.handlers[key]; ok {
		handlerFunc(c)
	}else{
		c.String( http.StatusNotFound , "")
	}
}
