package controller

import (
	"barrier-free-news/database"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	if r.RequestURI == "/home/detail" {
		t,_ := template.ParseFiles("detail.html")
		title,author,content := database.GetArticleByHref(r.FormValue("href"))
		var a article
		a.title = title
		a.author = author
		a.content = content
		t.Execute(w,a)
	}else if r.RequestURI == "/home/" {
		t,_ := template.ParseFiles("home.html")
		data := database.GetAllTitle()
		t.Execute(w,data)
	}
}

type article struct {
	title string
	author string
	content string
}

func clickDetail(w http.ResponseWriter,r *http.Request) {
	t,_ := template.ParseFiles("detail.html")
	title,author,content := database.GetArticleByHref(r.FormValue("href"))
	var a article
	a.title = title
	a.author = author
	a.content = content
	t.Execute(w,a)
}
