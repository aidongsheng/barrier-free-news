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





/***********************************************************************/
/***********************************************************************/
/*********************    简氏防务周刊爬虫部分   **************************/
/***********************************************************************/
/***********************************************************************/


func JanesDefenceWeekly() {
	c := colly.NewCollector()
	c.UserAgent = agent
	c.OnError(func(response *colly.Response, e error) {
		if e != nil {
			log.Fatal(e)
		}
	})
	c.OnRequest(func(request *colly.Request) {
		log.Printf("开始请求ID %d 请求URL %s",request.ID,request.URL)
	})

	c.OnHTML("div[class]", func(element *colly.HTMLElement) {
		ParseHtml.ParseJDWIndex(element)
	})
	c.OnScraped(func(response *colly.Response) {
		log.Print("抓取简氏防务周刊结束")
		janesDefenceWeeklyDetail(ParseHtml.DetailHrefs)
	})
	c.Visit("https://ihsmarkit.com/research-analysis/aerospace-defense-security.html")
}

func janesDefenceWeeklyDetail(hrefs []string) {
	c := colly.NewCollector()
	c.UserAgent = agent
	c.OnError(func(response *colly.Response, e error) {
		if e != nil {
			log.Fatal(e)
		}
	})
	c.OnRequest(func(request *colly.Request) {
		log.Printf("请求简氏防务详情ID %d 请求URL %s",request.ID,request.URL)
	})
	c.OnHTML("section[class]", func(element *colly.HTMLElement) {
		ParseHtml.ParseJDWDetail(element)
	})
	c.OnScraped(func(response *colly.Response) {
		log.Print("抓取简氏防务周刊详情页结束")
	})
	for _,href := range hrefs{
		c.Visit(href)
	}
}



/***********************************************************************/
/***********************************************************************/
/***********************    泰晤士报爬虫部分   ***************************/
/***********************************************************************/
/***********************************************************************/

func TheTimes() {
	c := colly.NewCollector()
	c.UserAgent = agent
	c.OnError(func(response *colly.Response, e error) {
		if e != nil {
			log.Fatal(e)
		}
	})
	c.OnRequest(func(request *colly.Request) {
		log.Printf("请求泰晤士报详情ID %d 请求URL %s",request.ID,request.URL)
	})
	c.OnHTML("h3[class]", func(element *colly.HTMLElement) {
		ParseHtml.ParseTimesTitle(element)
	})
	c.OnScraped(func(response *colly.Response) {
		log.Print("抓取简氏防务周刊详情页结束")
	})
	c.Visit("https://www.thetimes.co.uk/")
}

func TheTimesDetail([]string) {

}


/***********************************************************************/
/***********************************************************************/
/***********************  英国卫报爬虫部分   ***************************/
/***********************************************************************/
/***********************************************************************/

func GuardianIndex(){
	c := colly.NewCollector()
	c.UserAgent = agent
	c.OnError(func(response *colly.Response, e error) {
		if e != nil {
			log.Fatal(e)
		}
	})
	c.OnRequest(func(request *colly.Request) {
		log.Printf("请求英国卫报详情ID %d 请求URL %s",request.ID,request.URL)
	})
	c.OnHTML("a[class]", func(element *colly.HTMLElement) {
		ParseHtml.ParseGuardianTitle(element)
	})
	c.OnScraped(func(response *colly.Response) {
		log.Print("抓取英国卫报详情页结束")
	})
	c.Visit("https://www.thetimes.co.uk/")
}

/***********************************************************************/
/***********************************************************************/
/***********************  英国每日电讯报爬虫部分   ***************************/
/***********************************************************************/
/***********************************************************************/

func TelegraphIndex() {
	c := colly.NewCollector()
	c.OnError(func(response *colly.Response, e error) {
		log.Fatal(e)
	})
	c.OnRequest(func(request *colly.Request) {
		log.Printf("请求每日电讯报ID:%d 链接:%s",request.ID,request.URL.String())
	})
	c.OnHTML("h3[class]", func(element *colly.HTMLElement) {
		ParseHtml.ParseTelegraph(element)
	})
	c.OnScraped(func(response *colly.Response) {
		log.Print("抓取每日电讯报标题结束，开始抓取英国每日电讯报详情页")
		TelegraphDetail(ParseHtml.TelegraphDetailHrefs)
	})
	c.Visit("https://www.telegraph.co.uk/politics/")
}

func TelegraphDetail(hrefs []string) {
	c := colly.NewCollector()
	c.UserAgent = agent
	c.OnError(func(response *colly.Response, e error) {
		if e != nil {
			log.Fatal(e)
		}
	})
	c.OnRequest(func(request *colly.Request) {
		log.Printf("请求每日电讯报详情ID %d 请求URL %s",request.ID,request.URL)
	})
	c.OnHTML("div[class]", func(element *colly.HTMLElement) {
		ParseHtml.ParseTelegraphDetail(element)
	})
	c.OnScraped(func(response *colly.Response) {
		log.Print("抓取每日电讯报详情页结束")
	})
	for _,href := range hrefs{
		c.Visit(href)
	}
}
