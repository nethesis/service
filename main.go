package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nethesis/service/methods"
)

func main() {

	// init routers
	router := gin.Default()
	router.GET("/hello", methods.HelloWorld)

	// handle missing endpoint
	router.NoRoute(func(c *gin.Context) {
		c.Status(http.StatusNotFound)
	})

	router.Run()
}
