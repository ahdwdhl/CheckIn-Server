package utils

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

type Cookies map[string]string

var (
	ErrorUrl      error = errors.New("Url too short!")
	ErrorRequest  error = errors.New("request error!")
	ErrorResponse error = errors.New("response error!")
	ErrorCookie   error = errors.New("Cookie error!")
)

var globalClient *http.Client = nil
var globalCookieJar *cookiejar.Jar = nil
var request *http.Request = nil
var response *http.Response = nil
var Url *url.URL = nil
var err error
var jarBind bool = false
