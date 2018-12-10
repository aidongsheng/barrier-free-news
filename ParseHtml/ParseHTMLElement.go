package ParseHtml

import (
	"barrier-free-news/database"
	"barrier-free-news/translate"
	"github.com/gocolly/colly"
	"log"
)

/***********************************************************************/
/***********************************************************************/
/*********************    英国每日邮报解析部分   **************************/
/***********************************************************************/
/***********************************************************************/

/* 解析某一个单独的新闻标题 */
func dmArticleTitle(element *colly.HTMLElement) {
	element.ForEach("a[itemprop]", func(i int, element *colly.HTMLElement) {
		if element.Attr("itemprop") == "url" {
			link := element.Request.AbsoluteURL(element.Attr("href")	)	// 文章的链接
			text := element.Text				//	文章的标题
			transtext := translate.StartYoudaoFanyi(text)
			log.Printf("标题文字:%s 翻译文字:%s 链接:%s",text,transtext,link)
			database.InsertArticleList(text,transtext,link)
		}
	})
}

/* 解析 daily mail 首页数据 */
func DMIndex(element *colly.HTMLElement) {
	if element.Attr("class") == "cleared lead-alpha" {
		dmArticleTitle(element)	//	解析标题
	}
}

/* 解析 daily mail 详情页数据 */
func DMDetail(element *colly.HTMLElement) {

	if element.Attr("class") == "article-text wide  heading-tag-switch" {
		var title, author, content string
		element.ForEach("h2", func(i int, element *colly.HTMLElement) {
			if i == 0 {
				log.Printf("标题:%s",element.Text)
				title = element.Text
			}
		})

		element.ForEach("p", func(i int, element *colly.HTMLElement) {
			if element.Attr("class") == "author-section byline-plain" {
				element.ForEach("a", func(i int, element *colly.HTMLElement) {
					if i == 0 {
						log.Printf("作者: %s",element.Text)
						author = element.Text
					}
				})
			}
		})
		var tmpcontent string
		element.ForEach("p[class]", func(i int, element *colly.HTMLElement) {
			if element.Attr("class") == "mol-para-with-font" {
				tmpcontent = "<p>" + element.Text + "</p>"
				log.Printf("内容: %s",tmpcontent)
				content = content + tmpcontent
			}
		})
		database.InsertArticle(title,author,content,element.Response.Request.URL.String())
	}
}

/* 解析 daily mail 评论页数据 */
func ParseCommentEle(ele *colly.HTMLElement) {

}

/***********************************************************************/
/***********************************************************************/
/************************    印度时报解析部分   **************************/
/***********************************************************************/
/***********************************************************************/





/***********************************************************************/
/***********************************************************************/
/*********************    简氏防务周刊解析部分   **************************/
/***********************************************************************/
/***********************************************************************/

var DetailHrefs []string	//	待抓取详情页链接

func ParseJDWIndex(element *colly.HTMLElement) {
	if element.Attr("class") == "info image clearfix" {

		element.ForEach("header", func(i int, element *colly.HTMLElement) {
			var href string
			element.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
				href = element.Request.AbsoluteURL(element.Attr("href"))
				DetailHrefs = append(DetailHrefs,href)
			})

			timeAuthor := element.ChildText("span")
			title := element.ChildText("a")
			log.Printf("简氏防务周刊标题 %s 作者 %s 链接 %s",title,timeAuthor,href)
			database.InsertIntoJDW(title,timeAuthor,href)
		})
	}
}


func ParseJDWDetail(element *colly.HTMLElement) {

	if element.Attr("class") == "content basic-content bg-white no-border blog_post" {
		var title ,content string

		title = element.ChildText("h1")
		element.ForEach("span[class]", func(i int, element *colly.HTMLElement) {
			if element.Attr("class") == "blog_content" {
				element.ForEach("p", func(i int, element *colly.HTMLElement) {
					content = content + "<p>" + element.Text + "</p>"
				})
			}
		})
		database.InsertIntoJDWDetail(title,content,element.Request.URL.String())
	}
}