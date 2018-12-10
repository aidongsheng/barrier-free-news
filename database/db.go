package database

import (
	"database/sql"
	"log"
	"time"
)
import _ "github.com/go-sql-driver/mysql"


const dataSourceName  = "root:ads---@/wcc"


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


func InsertArticle(title string,author string,articleContent string,href string) {
	mydb,openErr := sql.Open("mysql",dataSourceName)
	if openErr != nil {
		log.Fatal(openErr)
	}

	defer mydb.Close()
	insertStmt, insertErr := mydb.Prepare("insert into article values (?,?,?,?,?)")
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
		_,resultErr := insertStmt.Exec(insertTime,title,author,articleContent,href)
		log.Printf("插入文章 %s",articleContent)
		if resultErr != nil{
			log.Fatal(resultErr)
		}
	}

}

/* 从数据库获取所有文章标题数据 */
func GetAllTitle() (map[string]string){
	mydb,openErr := sql.Open("mysql",dataSourceName)
	if openErr != nil {
		log.Fatal(openErr)
	}

	defer mydb.Close()
	rows,qErr := mydb.Query("select al_ti_trans,al_href from article_list")
	if qErr != nil {
		log.Fatal(qErr)
	}
	result := make(map[string]string)
	for rows.Next() {
		var href  string
		var titleTrans string
		if err := rows.Scan(&titleTrans,&href); err != nil {
			log.Fatal(err)
		}
		//href = "192.168.101.28:8000/home/detail?href=" + "dododo"
		if titleTrans != "" && href != ""{	//	防止出现空标题或空url
			result[href] = titleTrans
		}
	}
	return result
}

func GetArticleByHref(query string) (ti string,au string,hr string) {
	mydb,openErr := sql.Open("mysql",dataSourceName)
	if openErr != nil {
		log.Fatal(openErr)
	}

	defer mydb.Close()

	rows,qErr := mydb.Query("select * from article where href='" + query + "'")
	if qErr != nil{
		log.Fatal(qErr)
	}
	var title,author,content string
	for rows.Next() {
		err := rows.Scan(&title,&author,&content)
		if err != nil {
			log.Fatal(err)
		}
	}
	return title,author,content
}
/*插入简氏防务周刊*/
func InsertIntoJDW(title string, author string, href string) {
	mydb,openErr := sql.Open("mysql",dataSourceName)
	if openErr != nil {
		log.Fatal(openErr)
	}

	defer mydb.Close()

	stmt,err := mydb.Prepare("insert into janes_article values (?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	time := time.Now().Format("2006-01-02 15:04:05")
	_,ierr := stmt.Exec(time,title,author,href)
	if ierr != nil {
		log.Fatal(ierr)
	}
}
/*插入简氏防务周刊详情页数据*/
func InsertIntoJDWDetail(title string, content string, href string) {
	mydb,openErr := sql.Open("mysql",dataSourceName)
	if openErr != nil {
		log.Fatal(openErr)
	}

	defer mydb.Close()

	stmt,err := mydb.Prepare("insert into janes_detail values (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("插入简氏防务详情数据:%s %s %s",title,content,href)
	_,ierr := stmt.Exec(title,content,href)
	if ierr != nil {
		log.Fatal(ierr)
	}
}



/*插入简氏防务周刊详情页数据*/
func InsertIntoTelegraph(title string, titleTrans string, href string) {
	mydb,openErr := sql.Open("mysql",dataSourceName)
	if openErr != nil {
		log.Fatal(openErr)
	}

	defer mydb.Close()

	stmt,err := mydb.Prepare("insert into Telegraph values (?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("插入简氏防务详情数据:%s %s %s",title,titleTrans,href)
	strTime := time.Now().Format("2006-01-02 15:04:05")
	_,ierr := stmt.Exec(strTime,title,titleTrans,href)
	if ierr != nil {
		log.Fatal(ierr)
	}
}


/*插入每日电讯报详情页数据*/
func InsertIntoTelegraphDetail(title string, content string, href string) {
	mydb,openErr := sql.Open("mysql",dataSourceName)
	if openErr != nil {
		log.Fatal(openErr)
	}

	defer mydb.Close()

	stmt,err := mydb.Prepare("insert into telegraph_detail values (?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("插入每日电讯报详情数据:%s %s %s",title,content,href)
	strTime := time.Now().Format("2006-01-02 15:04:05")
	_,ierr := stmt.Exec(strTime,title,content,href)
	if ierr != nil {
		log.Fatal(ierr)
	}
}