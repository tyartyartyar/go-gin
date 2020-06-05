package main

import "errors"

type article struct {
	ID    int
	Title string
	Body  string
}

var articleList = []article{
	article{ID: 1, Title: "Title Satu", Body: "Body Satu"},
	article{ID: 2, Title: "Title Dua", Body: "Body Dua"},
}

func getAllArticles() []article {
	return articleList
}

func getArticleByID(ID int) (*article, error) {
	for _, v := range articleList {
		if v.ID == ID {
			return &v, nil
		}
	}
	return nil, errors.New("Article Not Found")
}
