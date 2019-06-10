package websever

import (
	"../utils"
	"github.com/axgle/mahonia"
	"time"
)

func NoCahce() {
	logger.Println(_url, " cookie 读取失败...")
	//请求url
	_, err := utils.GetUrlHtml(_url)
	if err != nil {
		logger.Fatal(err)
	}
	time.Sleep(100 * 5)
	logger.Println(_url, "打开主页.....")
	_, err = utils.PostUrlHtml(_url+"/auth/login", Pdatas)
	if err != nil {
		logger.Fatal(err)
	}
	time.Sleep(100 * 5)
	logger.Println(_url, "登录中.....")
}

func HasCache() string {
	body, err := utils.PostUrlHtml(_url+"/user/checkin", nil)
	if err != nil {
		logger.Fatal(err)
	}
	dec := mahonia.NewDecoder("gbk")
	body = dec.ConvertString(body)
	body = UnicodeToString(body)

	bodyHtml := "<body> " + body + "</body>"
	logger.Println(body)

	return bodyHtml
}