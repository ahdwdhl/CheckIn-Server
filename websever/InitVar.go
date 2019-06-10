package websever

import "container/list"

func InitGlobalValue() {
	Pdatas = make(map[string]string)
	Pdatas["email"] = "993534609@qq.com"
	Pdatas["passwd"] = "lyf110110"
	Pdatas["code"] = ""

	Urls[0] = "http://ssr.bingly.cn"
	_url = Urls[0]
	Urls[1] = "http://ssr.guopao.tk"
	Urls[2] = "https://coolji.ml"

	CookieFiles[0] = "bl.json"
	_ckFile = CookieFiles[0]
	CookieFiles[1] = "gp.json"
	CookieFiles[2] = "cj.json"
	Nodes = new(list.List)
}
