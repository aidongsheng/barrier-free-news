package Spiders

import (
	"barrier-free-news/ParseHtml"
	"barrier-free-news/database"
	"github.com/gocolly/colly"
	"log"
	"math/rand"
)

var (
	dmIndexUrl = "https://www.dailymail.co.uk"
	agent = []string{

		// Chrome
		//Win7:
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/14.0.835.163 Safari/535.1",
		// Firefox
		//Win7:
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:6.0) Gecko/20100101 Firefox/6.0",

		// Safari
		//Win7:
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534.50 (KHTML, like Gecko) Version/5.1 Safari/534.50",

		// Opera
		//Win7:
		"Opera/9.80 (Windows NT 6.1; U; zh-cn) Presto/2.9.168 Version/11.50",


		//Win7+ie9：
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 2.0.50727; SLCC2; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.3; .NET4.0C; Tablet PC 2.0; .NET4.0E)",

		//Win7+ie8：
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; InfoPath.3)",

		//WinXP+ie8：
		"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.1; Trident/4.0; GTB7.0)",

		//WinXP+ie7：
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 5.1)",

		//WinXP+ie6：
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1)",

		//傲游3.1.7在Win7+ie9,高速模式:
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; ) AppleWebKit/534.12 (KHTML, like Gecko) Maxthon/3.0 Safari/534.12",

		//傲游3.1.7在Win7+ie9,IE内核兼容模式:
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.3; .NET4.0C; .NET4.0E)",


		//搜狗3.0在Win7+ie9,IE内核兼容模式:
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.3; .NET4.0C; .NET4.0E; SE 2.X MetaSr 1.0)",

		//搜狗3.0在Win7+ie9,高速模式:
		"Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US) AppleWebKit/534.3 (KHTML, like Gecko) Chrome/6.0.472.33 Safari/534.3 SE 2.X MetaSr 1.0",


		//360浏览器3.0在Win7+ie9:
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.3; .NET4.0C; .NET4.0E)",

		//QQ浏览器6.9(11079)在Win7+ie9,极速模式:
		"Mozilla/5.0 (Windows NT 6.1) AppleWebKit/535.1 (KHTML, like Gecko) Chrome/13.0.782.41 Safari/535.1 QQBrowser/6.9.11079.201",

		//QQ浏览器6.9(11079)在Win7+ie9,IE内核兼容模式:
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/5.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.3; .NET4.0C; .NET4.0E) QQBrowser/6.9.11079.201",


		//阿云浏览器1.3.0.1724 Beta(编译日期2011-12-05)在Win7+ie9:
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; WOW64; Trident/5.0)",
		//	safari
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0.2 Safari/605.1.15",
		//	chrome
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36",
	}
	c *colly.Collector
)

/***********************************************************************/
/***********************************************************************/
/************************    每日邮报爬虫部分   **************************/
/***********************************************************************/
/***********************************************************************/
/* 爬取每日邮报首页文章数据 */
func DMCrawlIndex() {

	c = colly.NewCollector()

	c.UserAgent = agent[rand.Int()%len(agent)]

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
		log.Printf("结束抓取每日邮报首页 %s 开始抓取每日邮报详情页面 urls",response.Request.URL)

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

	c = colly.NewCollector()
	c.UserAgent = agent[rand.Int()%len(agent)]
	c.OnRequest(func(request *colly.Request) {
		log.Printf("抓取每日邮报详情ID: %d 链接: %s",request.ID,request.URL)
	})
	c.OnError(func(response *colly.Response, e error) {
		log.Printf("抓取每日邮报详情失败 %s\n失败原因 %s",response.Request.URL,e)
	})
	c.OnScraped(func(response *colly.Response) {
		log.Printf("结束抓取每日邮报详情 %s",response.Request.URL)
	})
	c.OnHTML("div[class]", func(element *colly.HTMLElement) {
		if element.Attr("class") == "article-text wide  heading-tag-switch" {
			log.Print("解析每日邮报详情页数据")
			ParseHtml.DMDetail(element)
		}
	})
	for _,href := range hrefs {
		c.Visit(href)
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
	c.UserAgent = agent[rand.Int()%len(agent)]
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
	c.UserAgent = agent[rand.Int()%len(agent)]
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
	c.UserAgent = agent[rand.Int()%len(agent)]
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
	c.IgnoreRobotsTxt = true
	coo := c.Cookies("https://www.thetimes.co.uk/")
	log.Printf("英国卫报cookie信息%s",coo)
	c.UserAgent = agent[rand.Int()%len(agent)]
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
	c.UserAgent = agent[rand.Int()%len(agent)]
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
