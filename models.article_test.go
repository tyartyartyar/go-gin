package main

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := articleList

	if len(alist) != len(articleList) {
		t.Fail()
	}

	for i, v := range alist {
		if v.ID != articleList[i].ID { // NOTE check that the ID is identical
			t.Fail()
			break
		}
	}

}
