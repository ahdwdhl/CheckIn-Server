package websever

import (
	"../utils"
	"container/list"
	"log"
)

var (
	Pdatas      utils.Cookies
	Urls        [3]string
	_url        string
	CookieFiles [3]string
	_ckFile     string
	logger      *log.Logger = nil
	Nodes       *list.List
)
