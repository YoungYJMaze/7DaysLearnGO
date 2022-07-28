package main

import (
	webServe "awesomeProject/http"
	"fmt"
	"net/http"
	"testing"
)

func main() {
	webServe.TestGetRouter(&testing.T{})
	engine := webServe.NewEngine()
	g1 := engine.Group("/youngyj")
	{
		g1.GET("/", func(c *webServe.Context) {
			c.HTML(http.StatusOK, "<h1>Hello</h1> ")
		})
		g1.GET("/hello/:name", func(c *webServe.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1> "+c.Param("name"))
		})
		g1.GET("/hello/", func(c *webServe.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1> ")
		})
	}
	engine.GET("/", func(c *webServe.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	engine.GET("/hello", func(context *webServe.Context) {
		context.String(http.StatusOK, "hello %s, you're at %s\n", context.Query("name"), context.Path)
	})
	engine.POST("/login", func(c *webServe.Context) {
		c.Json(http.StatusOK, map[string]string{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	engine.Run(":9999")
}

func youngyjHandler(writer http.ResponseWriter, request *http.Request) {
	for k, v := range request.Header {
		fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
	}
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}
