package cache

import (
	"../utils"
	"fmt"
	"github.com/json-iterator/go"
	"os"
)

var cookies utils.Cookies
var js = jsoniter.ConfigCompatibleWithStandardLibrary
var hasCookie bool = false

func InitCookies() {
	cookies, _ = utils.ReadEasyCookie()
}
func WtiteCookie(s string) {
	fp, err := os.OpenFile(s, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		os.Exit(-1)
	}
	defer fp.Close()
	hasCookie = true
	data, _ := js.Marshal(cookies)
	fp.Write(data)
}
func ReadCookie(s string) (bool, utils.Cookies) {
	fp, err := os.OpenFile(s, os.O_RDONLY, 0755)
	if err != nil {
		fmt.Println("read file error")
		hasCookie = false
		return hasCookie, nil
	}
	defer fp.Close()
	data := make([]byte, 512)
	n, _ := fp.Read(data)
	data = data[:n]
	if n <= 0 {
		fmt.Println("read json error")
		hasCookie = false
		return hasCookie, nil
	}
	cks := utils.Cookies{}
	_ = js.Unmarshal(data, &cks)

	return true, cks
}
