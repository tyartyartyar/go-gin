package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {

	articles := getAllArticles()

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}

func getArticle(c *gin.Context) {

	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		// fmt.Println(article)

		if article, err := getArticleByID(articleID); err == nil {
			c.HTML(
				http.StatusOK,
				"article.html",
				gin.H{
					"title":   article.Title,
					"payload": article,
				},
			)
		} else {
			c.AbortWithError(http.StatusNoContent, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
