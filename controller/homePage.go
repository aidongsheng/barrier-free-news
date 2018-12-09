package controller

import (
	"barrier-free-news/database"
	"database/sql"
	"html/template"
	"log"
	"net/http"
)

type ArticleTitle struct {
	time string			//	时间
	titleOri string			//	原标题
	titleTrans string		//	翻译标题
	href string				//	链接地址
}

/* 从数据库获取所有文章标题 */
func getAllTitle() (map[string]string){
	mydb,openErr := sql.Open("mysql",database.DataSourceName)
	if openErr != nil {
		log.Fatal(openErr)
	}

	defer mydb.Close()
	rows,qErr := mydb.Query("select al_title,al_ti_trans,al_href from article_list")
	if qErr != nil {
		log.Fatal(qErr)
	}
	result := make(map[string]string)
	for rows.Next() {
		rs := ArticleTitle{}
		if err := rows.Scan(&rs.titleOri,&rs.titleTrans,&rs.href); err != nil {
			log.Fatal(err)
		}
		result[rs.href] = rs.titleTrans
	}
	return result
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	t,_ := template.ParseFiles("home.html")
	data := getAllTitle()
	t.Execute(w,data)
}