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
	router := gin.Default()
	router.GET("/", home)
	router.GET("/:str", redirectPage)

	router.Run() // listen and serve on 0.0.0.0:8080
}

func home(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "Welcome to Url-shortener")
}

func redirectPage(c *gin.Context) {
	path := c.Request.URL.Path
	fmt.Println("The URL: ", c.Request.Host+c.Request.URL.Path)
	fmt.Println("Path : ", c.Request.URL.Path)
	fmt.Println("Redirecting to : ", pathsToUrls[path])
	if dest, ok := pathsToUrls[path]; ok {
		c.Redirect(302, dest)
	} else {
		c.Redirect(302, "http://127.0.0.1:8080")
	}

}
