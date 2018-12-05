package main

import (
	"awesomeProject1/translate"
)

func main() {
	translate.Youdao()
	//c := colly.NewCollector()
	//c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36"
	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting ", r.URL)
	//})
	//c.OnError(func(r *colly.Response, e error) {
	//	log.Println("Something went wrong:", e)
	//})
	//c.OnResponse(func(r *colly.Response) {
	//	fmt.Println("Visited ", r.Request.URL)
	//})
	//c.OnHTML("h2", func(e *colly.HTMLElement) {
	//	print(e.Text)
	//})
	//c.OnHTML("tr td:nth-of-type(1)", func(e *colly.HTMLElement) {
	//	fmt.Println("First column of a table row:", e.Text)
	//})
	//c.OnXML("//h1", func(e *colly.XMLElement) {
	//	fmt.Println(e.Text)
	//})
	//c.OnScraped(func(r *colly.Response) {
	//	fmt.Println("Finished ",r.Request.URL)
	//})
	//fmt.Printf("user-agent ",c.UserAgent)
	//c.Visit("https://www.dailymail.co.uk")

}
//
//package main
//
//import (
//	"fmt"
//	"log"
//	"net/http"
//
//	"github.com/PuerkitoBio/goquery"
//)
//
//func ExampleScrape() {
//	// Request the HTML page.
//	res, err := http.Get("http://metalsucks.net")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer res.Body.Close()
//	if res.StatusCode != 200 {
//		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
//	}
//
//	// Load the HTML document
//	doc, err := goquery.NewDocumentFromReader(res.Body)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// Find the review items
//	doc.Find("article").Each(func(i int, s *goquery.Selection) {
//		// For each item found, get the band and title
//		band := s.Find("a").Text()
//		//title := s.Find("i").Text()
//		fmt.Printf("%s",band)
//	})
//}
//
//func main() {
//	ExampleScrape()
//}