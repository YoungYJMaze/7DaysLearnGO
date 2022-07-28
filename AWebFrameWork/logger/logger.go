package logger

import (
	webServe "awesomeProject/http"
	"log"
	"time"
)

func Logger() webServe.HandlerFunc {
	return func(c *webServe.Context) {
		t := time.Now()
		// process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v ", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
