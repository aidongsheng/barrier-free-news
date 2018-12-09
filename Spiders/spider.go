package Spiders

import (
	"barrier-free-news/ParseHtml"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
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

	c.OnError(func(response *colly.Response, e error) {
		log.Fatal(e)
	})

	c.OnRequest(func(request *colly.Request) {
		log.Printf("请求ID: %d 链接: %s",request.ID,request.URL)
	})

	c.OnError(func(response *colly.Response, e error) {
		log.Printf("抓取失败信息: %s 请求ID: %s 链接: %s",e,response.Request.ID,response.Request.URL)
	})
	//	新闻标题页面解析
	c.OnHTML("div[class]", func(element *colly.HTMLElement) {
		ParseHtml.DMIndex(element)

		if element.Attr("class") == "cleared lead-alpha" {
			element.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
				c.Visit(element.Request.AbsoluteURL(element.Attr("href")))
			})
		}
	})

	//	新闻详情页解析
	c.OnHTML("div[class]", func(element *colly.HTMLElement) {
		ParseHtml.DMDetail(element)
	})

	c.OnScraped(func(response *colly.Response) {
		log.Printf("结束抓取 %s %s",response.Request.URL,c.String())
	})

	c.Visit(dmIndexUrl)
}

/* 爬取每日邮报文章详情页数据 */
func DMCrawlDetail(hrefs []string) {

	if c1 == nil {
		c1 = colly.NewCollector()
	}

	q,_  := queue.New(
		2,
		&queue.InMemoryQueueStorage{MaxSize:10000},
	)

	for _,url := range hrefs{
		q.AddURL(url)
	}

	c1.UserAgent = agent
	c1.OnError(func(response *colly.Response, e error) {
		log.Printf("c1 抓取失败 %s\n失败原因 %s",response.Request.URL,e)
	})
	c1.OnRequest(func(request *colly.Request) {
		log.Printf("c1 开始抓取 %s",request.URL)
	})
	c1.OnScraped(func(response *colly.Response) {
		log.Printf("c1 结束抓取 %s",response.Request.URL)
	})
	c1.OnHTML("div[class=(article-text wide  heading-tag-switch)]", func(element *colly.HTMLElement) {
		log.Printf("c1 解析 HTML 结果 %s",element.ChildText("h2"))
	})

	q.Run(c)

}

/* 爬取每日邮报文章评论数据 */
func DMCrawlComment() {

}

/***********************************************************************/
/***********************************************************************/
/************************    印度时报爬虫部分   **************************/
/***********************************************************************/
/***********************************************************************/