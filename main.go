package main

import (
	"barrier-free-news/controller"
	"net/http"
)

func main() {
	//Spiders.DMCrawlIndex()

	srv := http.Server{
		Addr:"192.168.1.3:8000",
	}
	http.HandleFunc("/home/",controller.HomePage)
	//http.HandleFunc("/home/:detail",controller.Detail)
	srv.ListenAndServe()
}
