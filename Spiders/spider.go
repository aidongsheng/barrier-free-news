package Spiders

import (
	"barrier-free-news/ParseHtml"
	"barrier-free-news/database"
	"github.com/gocolly/colly"
	"log"
)

var (
	dmIndexUrl = "https://www.dailymail.co.uk"
	agent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36"
	c *colly.Collector
	c1 *colly.Collector
)


/***********************************************************************/
/***********************************************************************/
/************************    每日邮报爬虫部分   **************************/
/***********************************************************************/
/***********************************************************************/
/* 爬取每日邮报首页文章数据 */
func DMCrawlIndex() {

	c = colly.NewCollector()

	c.UserAgent = agent

	c.OnRequest(func(request *colly.Request) {
		log.Printf("请求ID: %d 链接: %s",request.ID,request.URL)
	})

	c.OnError(func(response *colly.Response, e error) {
		log.Fatal(e)
		log.Printf("抓取失败信息: %s 请求ID: %s 链接: %s",e,response.Request.ID,response.Request.URL)
	})
	//	新闻标题页面解析
	c.OnHTML("div[class]", func(element *colly.HTMLElement) {
		if element.Attr("class") == "cleared lead-alpha" {
			ParseHtml.DMIndex(element)
			element.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
				href := element.Attr("href")
				if len(href) > 4 && href[:3] == "http" {
					c.Visit(element.Request.AbsoluteURL(element.Attr("href")))
				}
			})
		}
		if element.Attr("class") == "article-text wide  heading-tag-switch" {
			ParseHtml.DMDetail(element)
		}
	})


	c.OnScraped(func(response *colly.Response) {
		log.Printf("结束抓取 %s %s",response.Request.URL,c.String())

		//	开始详情页新闻抓取
		hrefs := []string{}
		for key,_ := range database.GetAllTitle() {
			hrefs = append(hrefs, key)
		}
		DMCrawlDetail(hrefs)
	})

	c.Visit(dmIndexUrl)
}

/* 爬取每日邮报文章详情页数据 */
func DMCrawlDetail(hrefs []string) {

	c1 = colly.NewCollector()
	c1.UserAgent = agent
	c1.OnRequest(func(request *colly.Request) {
		log.Printf("请求ID: %d 链接: %s",request.ID,request.URL)
	})
	c1.OnError(func(response *colly.Response, e error) {
		log.Printf("c1 抓取失败 %s\n失败原因 %s",response.Request.URL,e)
	})
	c1.OnRequest(func(request *colly.Request) {
		log.Printf("c1 开始抓取 %s",request.URL)
	})
	c1.OnScraped(func(response *colly.Response) {
		log.Printf("c1 结束抓取 %s",response.Request.URL)
	})
	c1.OnHTML("div[class]", func(element *colly.HTMLElement) {
		if element.Attr("class") == "article-text wide  heading-tag-switch" {
			log.Print("解析详情页数据")
			ParseHtml.DMDetail(element)
		}
	})
	for _,href := range hrefs {
		c1.Visit(href)
	}
}

/* 爬取每日邮报文章评论数据 */
func DMCrawlComment() {

}

/***********************************************************************/
/***********************************************************************/
/************************    印度时报爬虫部分   **************************/
/***********************************************************************/
/***********************************************************************/