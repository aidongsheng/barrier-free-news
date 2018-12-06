package translate

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

type baiduResult struct {
	From string `json:"from"`
	To   string	`json:"to"`
	Trans_result [](map[string]string) `json:"trans_result"`
}

type youdaoResult struct {
	TSpeakUrl string `json:"tSpeakUrl"`
	Query   string	`json:"query"`
	Translation []string `json:"translation"`
	ErrorCode string	`json:"error_code"`
	Dict map[string]string	`json:"dict"`
	Webdict map[string]string `json:"webdict"`
	L string `json:"l"`
	SpeakUrl string `json:"speak_url"`
}

/* 生成 sign 值 */
func genSign(appid string,q string , salt string, key string) string {
	h := md5.New()
	h.Write([]byte(appid + q + salt + key))
	return hex.EncodeToString(h.Sum(nil))
}

/* 有道翻译相关参数 */
var (
	strYoudaoAPI = "http://openapi.youdao.com/api"
	strYoudaoQ string
	strYoudaoAppId = "641ea2ca7f606210"
	strYoudaoSalt = string(rand.Int())
	strYoudaoKey = "cChiAbqDUAVZOsWJzip9QtNL8NtVnY1i"
	mapYoudaoParameter = make(map[string][]string)
)
/* 开始有道翻译 */
func StartYoudaoFanyi(fromString string) string {

	strYoudaoQ = fromString
	strYoudaoSign := genSign(strYoudaoAppId,strYoudaoQ,strYoudaoSalt,strYoudaoKey)

	mapYoudaoParameter["q"] = []string{strYoudaoQ}
	mapYoudaoParameter["from"] = []string{"EN"}
	mapYoudaoParameter["to"] = []string{"zh-CHS"}
	mapYoudaoParameter["appKey"] = []string{strYoudaoAppId}
	mapYoudaoParameter["salt"] = []string{strYoudaoSalt}
	mapYoudaoParameter["sign"] = []string{strYoudaoSign}

	res,err := http.PostForm(strYoudaoAPI,mapYoudaoParameter)
	if err != nil {
		log.Fatal(err)
	}

	data,e := ioutil.ReadAll(res.Body)
	if e != nil {
		log.Fatal(e)
	}

	var ydResult youdaoResult
	json.Unmarshal(data,&ydResult)
	return string(ydResult.Translation[0])
}

/* 百度翻译相关参数 */
var (
	strBaiduAPI = "http://api.fanyi.baidu.com/api/trans/vip/translate"
	strBaiduQ  string
	strBaiduFrom = "en"
	strBaiduTo = "zh"
	strBaiduAppId = "20181206000244467"
	strBaiduKey = "7IwStgeiPlJR4ScGKl0n"
	strBaiduSalt = string(rand.Int())
	mapBaiduParameter = make(map[string][]string)
)

func StartBaiduFanyi(fromString string) string {
	strBaiduQ = fromString
	strBaiduSign := genSign(strBaiduAppId,strBaiduQ,strBaiduSalt,strBaiduKey)

	mapBaiduParameter["q"] = []string{strBaiduQ}
	mapBaiduParameter["from"] = []string{strBaiduFrom}
	mapBaiduParameter["to"] = []string{strBaiduTo}
	mapBaiduParameter["appid"] = []string{strBaiduAppId}
	mapBaiduParameter["salt"] = []string{strBaiduSalt}
	mapBaiduParameter["sign"] = []string{strBaiduSign}
	res,err := http.PostForm(strBaiduAPI,mapBaiduParameter)
	if err != nil {
		log.Fatal(err)
	}
	data,readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		 log.Fatal(readErr)
	}
	log.Print(string(data))
	var a baiduResult
	json.Unmarshal(data,&a)

	return string(data)
}
