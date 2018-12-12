package controller

import (
	"barrier-free-news/database"
	"fmt"
	"html/template"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {

	t,_ := template.ParseFiles("home.html")
	data := database.GetAllTitle()
	t.Execute(w,data)
}

type article struct {
	title string
	author string
	content string
}

func ClickDetail(w http.ResponseWriter,r *http.Request) {
	fmt.Fprintf(w,"hello")
}
