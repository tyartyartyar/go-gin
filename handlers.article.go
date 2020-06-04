package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	fmt.Println("showIndexPage")
	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title": "HomePage",
		},
	)
}
