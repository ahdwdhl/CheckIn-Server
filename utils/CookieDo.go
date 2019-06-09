package utils

import (
	"net/http"
	"net/http/cookiejar"
)

func ReadEasyCookie() (Cookies, error) {
	if !jarBind || Url == nil {
		return nil, ErrorCookie
	}
	gCurCookies := globalCookieJar.Cookies(Url)
	cks := make(Cookies)
	for i := 0; i < len(gCurCookies); i++ {
		cur := gCurCookies[i]
		cks[cur.Name] = cur.Value
	}
	return cks, nil
}

func SetCookieJar(c Cookies) {

	if globalCookieJar == nil {
		globalCookieJar, _ = cookiejar.New(nil)
	}
	var cookies []*http.Cookie
	for k, v := range c {
		curC := new(http.Cookie)
		curC.Name = k
		curC.Value = v
		cookies = append(cookies, curC)
	}
	globalCookieJar.SetCookies(Url, cookies)
}
