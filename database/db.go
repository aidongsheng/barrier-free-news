package database

import (
	"database/sql"
	"log"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

const dataSourceName  = "root:ads---@/wcc"

func InsertArticle(title string , title_translated string, href string,imgs string)  {

	if len(title) > 1000 {
		title = title[:1000]
	}

	mydb,openErr := sql.Open("mysql",dataSourceName)
	if openErr != nil {
		log.Fatal(openErr)
	}

	defer mydb.Close()

	insertStmt, insertErr := mydb.Prepare("insert into article_list values (?,?,?,?,?)")
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
		//insertTime
		//timeStr := time.Now().Format("2006-01-02 15:04:05")
		_,resultErr := insertStmt.Exec(insertTime,title,title_translated,href,imgs)
		log.Printf("插入数据:%s %s %s %s",insertTime,title,title_translated,href)
		if resultErr != nil{
			log.Fatal(resultErr)
		}
	}
}