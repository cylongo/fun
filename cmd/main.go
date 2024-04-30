package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.StaticFile("/", "./static/index.html")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "pong")
	})
	err := r.Run()
	if err != nil {
		return
	}
}
