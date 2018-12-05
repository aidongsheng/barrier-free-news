package translate

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

var q = "George HW Bush once worried that his funeral would be a non-event, a fear that proved unfounded today as thousands lined up to say goodbye to the 41st president of the United States. As guests streamed into the cathedral, the Bush family gathered at the front of the US Capitol, where Bush's body has been lying in state in the rotunda for three days. The family stood in a line as Bush Senior's casket was led out by the US Air Force Honor Guard. George W Bush was clearly emotional and appeared to be holding back tears as he watched his father's body taken down the steps and loaded into the presidential hearse."
var appkey = "641ea2ca7f606210"
var salt = string(rand.Int())
var key = "cChiAbqDUAVZOsWJzip9QtNL8NtVnY1i"
func Youdao() {

	h := md5.New()
	h.Write([]byte(appkey + q + salt + key))
	b := hex.EncodeToString(h.Sum(nil))
	var para = make(map[string][]string)

	para["q"] = []string{q}
	para["from"] = []string{"EN"}
	para["to"] = []string{"zh-CHS"}
	para["appKey"] = []string{appkey}
	para["salt"] = []string{salt}
	para["sign"] = []string{string(b)}

	res,err := http.PostForm("http://openapi.youdao.com/api",para)
	if err != nil {
		log.Fatal(err)
	}

	data,e := ioutil.ReadAll(res.Body)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("%s",string(data))
}
