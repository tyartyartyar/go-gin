package main

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
