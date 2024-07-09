package main

import (
	"net/http"

	"biubiu"
)

func main() {
	r := biubiu.New()
	r.GET("/", func(c *biubiu.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *biubiu.Context) {
		// expect /hello?name=biubiubiu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *biubiu.Context) {
		c.JSON(http.StatusOK, biubiu.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
