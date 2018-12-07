package database

import (
	"database/sql"
	"log"
	"time"
)
import _ "github.com/go-sql-driver/mysql"


const dataSourceName  = "root:ads19890823/@@/wcc"


/***********************************************************************/
/***********************************************************************/
/*********************    英国每日邮报数据库部分   ************************/
/***********************************************************************/
/***********************************************************************/

/* 插入标题表格 */
func InsertArticleList(title string , title_translated string, href string)  {

	if len(title) > 1000 {
		title = title[:1000]
	}

	mydb,openErr := sql.Open("mysql",dataSourceName)
	if openErr != nil {
		log.Fatal(openErr)
	}

	defer mydb.Close()

	insertStmt, insertErr := mydb.Prepare("insert into article_list values (?,?,?,?)")
	if insertErr != nil {
		log.Fatal(insertErr)
	}else {
		y := time.Now().Year()
		m := time.Now().Month()
		d := time.Now().Day()
		h := time.Now().Hour()
		min := time.Now().Minute()
		sec := time.Now().Second()
		nano := time.Now().Nanosecond()
		insertTime := time.Date(y,m,d,h,min,sec,nano,time.Local)

		_,resultErr := insertStmt.Exec(insertTime,title,title_translated,href)
		log.Printf("插入数据:%s %s %s %s",insertTime,title,title_translated,href)
		if resultErr != nil{
			log.Fatal(resultErr)
		}
	}
}
/* 从数据库获取文章标题链接 */
func GetArticleTitleHrefs() ([]string){
	mydb,openErr := sql.Open("mysql",dataSourceName)
	if openErr != nil {
		log.Fatal(openErr)
	}

	defer mydb.Close()
	rows,qErr := mydb.Query("select al_href from article_list")
	if qErr != nil {
		log.Fatal(qErr)
	}
	var arrHref = make([]string,1000)
	var i = 0
	for rows.Next() {
		var href string
		if err := rows.Scan(&href); err != nil {
			log.Fatal(err)
		}
		log.Printf("从数据库中取得的链接是 %s",href)
		arrHref[i] = href
		i++
	}
	return arrHref
}


func InsertArticle(title string,author string,articleContent string) {
	mydb,openErr := sql.Open("mysql",dataSourceName)
	if openErr != nil {
		log.Fatal(openErr)
	}

	defer mydb.Close()
	insertStmt, insertErr := mydb.Prepare("insert into article values (?,?,?,?)")
	if insertErr != nil {
		log.Fatal(insertErr)
	}else {
		y := time.Now().Year()
		m := time.Now().Month()
		d := time.Now().Day()
		h := time.Now().Hour()
		min := time.Now().Minute()
		sec := time.Now().Second()
		nano := time.Now().Nanosecond()
		insertTime := time.Date(y,m,d,h,min,sec,nano,time.Local)
		_,resultErr := insertStmt.Exec(insertTime,title,author,articleContent)
		log.Printf("插入文章 %s",articleContent)
		if resultErr != nil{
			log.Fatal(resultErr)
		}
	}

}