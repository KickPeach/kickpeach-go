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

	v1 := r.Group("/v1")

	{
		v1.GET("/", func(c *kickpeach.Context) {
			c.HTML(http.StatusOK, "<h1>Hello kickpeach</h1>")
		})

		v1.GET("/hello", func(c *kickpeach.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *kickpeach.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *kickpeach.Context) {
			c.JSON(http.StatusOK, kickpeach.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
