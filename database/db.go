package database

import (
	"database/sql"
	"log"
	"time"
)
import _ "github.com/go-sql-driver/mysql"


const DataSourceName  = "root:ads---@/wcc"


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

	mydb,openErr := sql.Open("mysql",DataSourceName)
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


func InsertArticle(title string,author string,articleContent string) {
	mydb,openErr := sql.Open("mysql",DataSourceName)
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


