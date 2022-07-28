package webServe

import (
	"log"
	"net/http"
	"strings"
)

// Router module
type router struct {
	handlers map[string]HandlerFunc
	route    map[string]*node
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc), route: make(map[string]*node)}
}

// add a router  make something fundamental for get/post
func (router *router) addRoute(method, pattern string, handlerFunc HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	_, ok := router.route[method]
	parts := parsePattern(pattern)
	if !ok {
		router.route[method] = &node{}
	}
	router.route[method].insert(pattern, parts, 0)
	router.handlers[key] = handlerFunc
}

func (router *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	route, ok := router.route[method]
	if !ok {
		return nil, nil
	}
	n := route.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			// 如果为 : 符 则使用 实际参数替换  if find :  use real param instead
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			// 如果为 * 符  则将之后的所有路径都加入到当前的 * 之后 if find * add remained part to params
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	// not found
	return nil, nil
}

func (group *RouterGroup) addRouter(method string, pattern string, handlerFunc HandlerFunc) {
	gPattern := group.prefix + pattern
	group.engine.router.addRoute(method, gPattern, handlerFunc)

}

func (router *router) handle(c *Context) {
	n, params := router.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers, router.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s \n", c.Path)
		})
	}
	c.Next()
}

func parsePattern(pattern string) (res []string) {
	temp := strings.Split(pattern, "/")
	res = make([]string, 0)
	for _, x := range temp {
		if x != "" {
			res = append(res, x)
			if x[0] == '*' {
				break
			}
		}
	}
	return
}
