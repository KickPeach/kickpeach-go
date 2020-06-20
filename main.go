package main

import (
	"kickpeach"
	"net/http"
)

func main() {
	r := kickpeach.New()
	r.GET("/", func(c *kickpeach.Context) {
		c.HTML(http.StatusOK,"<h1>Hello Kickpeach</h1>")
	})

	r.GET("/hello", func(c *kickpeach.Context) {
		c.String(http.StatusOK,"hello %s,you are at %s\n",c.Query("name"),c.Path)
	})


	r.POST("/login", func(c *kickpeach.Context) {
		c.JSON(http.StatusOK,kickpeach.H{
			"username":c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})


	r.Run(":9999")
}