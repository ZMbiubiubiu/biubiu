package main

import (
	"biubiu"
)

func main() {
	r := biubiu.New()
	r.GET("/hi", func(c *biubiu.Context) {
		c.String(200, "URL.Path = %q\n", c.Path)
	})

	r.GET("/hello", func(c *biubiu.Context) {
		c.JSON(200, map[string]any{"name": "Tom", "age": 12})
	})
	r.GET("/p/:lang/doc", func(c *biubiu.Context) {
		c.JSON(200, map[string]any{"lang": c.Params["lang"], "doc": "https://" + c.Params["lang"] + ".dev/doc"})
	})
	r.POST("/p/go/src", func(c *biubiu.Context) {
		c.JSON(200, map[string]any{"lang": "go", "src": "https://github.com/golang/go/tree/master/src"})
	})
	r.Run(":9999")
}
