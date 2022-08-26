package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var pathsToUrls = map[string]string{
	"/chaimettoast": "https://www.youtube.com/watch?v=LMEdbBK4bk0",
	"/das":          "https://github.com/athul-dev-7",
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
