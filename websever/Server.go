package websever

import (
	"../cache"
	"../utils"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func IndexHtml(w http.ResponseWriter, r *http.Request) {
	//_ = r.ParseForm()
	//datas := make(map[string]string)
	//for k, v := range r.Form {
	//	datas[k] = strings.Join(v, "")
	//}
	var res string
	res = "<body><h1>hello world</h1></body>"
	_, _ = fmt.Fprintf(w, res)
}
func UnicodeToString(s string) string {
	Unicode := strings.Split(s, "\\u")
	var res string
	for _, v := range Unicode {
		if len(v) < 1 {
			continue
		}
		temp, err := strconv.ParseInt(v, 16, 32)
		if err != nil {
			continue
		}
		res += fmt.Sprintf("%c", temp)
	}
	return res
}
func Check(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	Rdatas := make(map[string]string)
	for k, v := range r.Form {
		Rdatas[k] = strings.Join(v, "")
	}
	Case := Rdatas["s"][0] - '1'
	if Case >= 3 || Case < 0 {
		return
	}
	_url, _ckFile = Urls[Case], CookieFiles[Case]

	_h, r_cookie := cache.ReadCookie(_ckFile)
	if !_h || r_cookie == nil {
		NoCahce()
		bodyHtml := HasCache()
		cache.InitCookies()
		cache.WtiteCookie(_ckFile)
		logger.Println(_url, "cookie 写入成功..")
		_, _ = fmt.Fprintf(w, bodyHtml)
	} else {
		_, _ = utils.GetUrlHtml(_url + "/user")
		time.Sleep(100 * 5)
		utils.SetCookieJar(r_cookie)
		logger.Println(_url, "cookie 设置成功.....")

		bodyHtml := HasCache()
		_, _ = fmt.Fprintf(w, bodyHtml)
	}

}

func BindAPI() {
	http.HandleFunc("/", IndexHtml)
	http.HandleFunc("/check", Check)
}

func Start(g *log.Logger) {
	BindAPI()
	logger = g
	InitGlobalValue()
	_ = http.ListenAndServe(":8080", nil)
}
