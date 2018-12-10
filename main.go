package main

import "barrier-free-news/Spiders"

func main() {
	//Spiders.DMCrawlIndex()
	//
	//srv := http.Server{
	//	Addr:"192.168.101.28:8000",
	//}
	//http.HandleFunc("/home/",controller.HomePage)
	//http.HandleFunc("/home/detail/",controller.ClickDetail)
	//srv.ListenAndServe()
	Spiders.JanesDefenceWeekly()
}
