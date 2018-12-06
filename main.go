package main

import (
	"barrier-free-news/translate"
	"github.com/gocolly/colly"
	"log"
)

func main() {

	log.Print(translate.StartYoudaoFanyi("i'm han meimei"))
	translate.StartBaiduFanyi("Why Link My Offer To Decision On My Extradition")

	c := colly.NewCollector()
	c.MaxDepth = 2
	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		c.Visit(element.Attr("href"))
		log.Print(translate.StartBaiduFanyi(element.Text))
	})
	c.OnError(func(response *colly.Response, e error) {
		log.Fatal(e)
	})
	c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36"
	//c.MaxDepth = 0
	//c.MaxBodySize = 0
	//c.CacheDir = "/Users/josan/Documents"
	//c.IgnoreRobotsTxt = true
	c.Visit("https://www.ndtv.com/")
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