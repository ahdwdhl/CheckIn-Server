package websever

import (
	"../cache"
	"../utils"
	"fmt"
	"github.com/axgle/mahonia"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var logger *log.Logger = nil
var (
	url    string
	ckFile string
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
	Case := Rdatas["s"]
	if Case == "1" {
		url = "http://ssr.bingly.cn"
		ckFile = "bl.json"
	} else {
		url = "http://ssr.guopao.tk"
		ckFile = "gp.json"
	}
	Pdatas := make(map[string]string)
	Pdatas["email"] = "993534609@qq.com"
	Pdatas["passwd"] = "lyf110110"
	Pdatas["code"] = ""

	_h, r_cookie := cache.ReadCookie(ckFile)

	if !_h || r_cookie == nil {
		logger.Println(url, " cookie 读取失败...")
		//请求url
		body, _ := utils.GetUrlHtml(url)
		time.Sleep(100 * 5)
		logger.Println(url, "打开主页.....")
		body, _ = utils.PostUrlHtml(url+"/auth/login", Pdatas)
		time.Sleep(100 * 5)
		logger.Println(url, "登录中.....")

		body, _ = utils.PostUrlHtml(url+"/user/checkin", nil)
		dec := mahonia.NewDecoder("gbk")
		body = dec.ConvertString(body)
		body = UnicodeToString(body)

		bodyHtml := "<body> " + body + "</body>"
		logger.Println(body)

		cache.InitCookies()
		cache.WtiteCookie(ckFile)
		logger.Println(url, "cookie 写入成功..")
		fmt.Fprintf(w, bodyHtml)
	} else {
		_, _ = utils.GetUrlHtml(url + "/user/checkin")
		time.Sleep(100 * 5)
		utils.SetCookieJar(r_cookie)

		logger.Println(url, "cookie 设置成功.....")

		body, _ := utils.PostUrlHtml(url+"/user/checkin", nil)
		dec := mahonia.NewDecoder("gbk")
		body = dec.ConvertString(body)
		body = UnicodeToString(body)
		bodyHtml := "<body> " + body + "</body>"
		logger.Println(body)
		fmt.Fprintf(w, bodyHtml)
	}

}

func BindAPI() {
	http.HandleFunc("/", IndexHtml)
	http.HandleFunc("/check", Check)
}

func Start(g *log.Logger) {
	BindAPI()
	logger = g
	_ = http.ListenAndServe(":8080", nil)
}
