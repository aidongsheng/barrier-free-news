package Spiders

import (
	"github.com/gocolly/colly"
	"log"
)

var (
	dmIndexUrl = "https://www.dailymail.co.uk"
)

/* 爬取每日邮报首页文章数据 */
func DMCrawlIndex() {
	c := colly.NewCollector()
	c.OnRequest(func(request *colly.Request) {
		log.Printf("开始抓取 第 %d %s",request.ID,request.URL)
	})
	c.MaxDepth = 1
	c.OnError(func(response *colly.Response, e error) {
		log.Printf("抓取 %s 失败 %s",response.Request.URL,e)
	})
	c.OnHTML("div[class] a[itemprop]", func(element *colly.HTMLElement) {
		//title := string(element.Text)
		//imgs := element.ChildAttrs("img","src")
		//var imgArr string
		//for img := range imgs {
		//	imgArr = imgArr + ","+ string(img)
		//}
		href := element.Attr("href")
		if element.Attr("href")[:3] != "http" {
			href = dmIndexUrl + element.Attr("href")
			c.Visit(href)
		}else {
			c.Visit(href)
		}

		//translatedTitle := translate.StartBaiduFanyi(element.Text)
		//database.InsertArticle(title,translatedTitle,href,imgArr)

	})
	c.OnHTML("div[class=(article-text wide  heading-tag-switch)]", func(element *colly.HTMLElement) {
		log.Printf("%s",element.Attr("p"))
	})
	c.OnScraped(func(response *colly.Response) {
		log.Printf("结束抓取 %s %s",response.Request.URL,c.String())
	})
	c.OnError(func(response *colly.Response, e error) {
		log.Fatal(e)
	})
	c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36"
	c.Visit(dmIndexUrl)
}

/* 爬取每日邮报文章详情页数据 */
func DMCrawlDetail() {

}

/* 爬取每日邮报文章评论数据 */
func DMCrawlComment() {

}