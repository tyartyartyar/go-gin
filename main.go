package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	router = gin.Default()

	router.LoadHTMLGlob("template/*")

	router.GET("/", func(c *gin.Context) {
		fmt.Println(c)

		c.HTML(
			http.StatusOK,
			"index.html",
			gin.H{
				"title": "Home Page",
			},
		)

	})

	router.Run(":9000")

}
