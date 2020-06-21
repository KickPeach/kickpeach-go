package main

import (
	"kickpeach"
	"net/http"
)

func main() {
	r := kickpeach.New()
	r.GET("/", func(c *kickpeach.Context) {
		c.HTML(http.StatusOK, "<h1>Hello kickpeach</h1>")
	})

	r.GET("/hello", func(c *kickpeach.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *kickpeach.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *kickpeach.Context) {
		c.JSON(http.StatusOK, kickpeach.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}