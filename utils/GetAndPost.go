package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func InitGlobalVar() {
	globalClient = nil
	globalClient = nil
	globalCookieJar = nil
	request = nil
	response = nil
	Url = nil
	jarBind = false
}

func GetUrlHtml(_url string) (string, error) {
	if len(_url) <= 0 {
		fmt.Println(_url, "....")
		return "", ErrorUrl
	}
	if globalClient == nil {
		globalClient = new(http.Client)
	}
	if globalCookieJar == nil {
		globalCookieJar, _ = cookiejar.New(nil)
	}
	if !jarBind {
		globalClient.Jar = globalCookieJar
		jarBind = true
	}

	request, err = http.NewRequest("GET", _url, nil)
	if err != nil {
		fmt.Println("requset ...")
		return "", ErrorRequest
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")

	response, err = globalClient.Do(request)
	if err != nil {
		fmt.Println("response ...")
		return "", ErrorResponse
	}

	defer response.Body.Close()
	var bodyBytes []byte
	bodyBytes, _ = ioutil.ReadAll(response.Body)

	var bodyString string
	bodyString = fmt.Sprintf("%s", bodyBytes)
	Url = request.URL
	//enc := mahonia.NewDecoder("GB18030")
	return bodyString, nil
}
func EncodeDatas(datas map[string]string) []byte {
	postValue := url.Values{}
	for key, value := range datas {
		postValue.Set(key, value)
	}
	encode_data := postValue.Encode()
	postByte := []byte(encode_data)
	return postByte
}
func PostUrlHtml(_url string, datas map[string]string) (string, error) {
	if len(_url) <= 0 {
		fmt.Println(_url, "....")
		return "", ErrorUrl
	}
	if globalClient == nil {
		globalClient = new(http.Client)
	}
	if globalCookieJar == nil {
		globalCookieJar, _ = cookiejar.New(nil)
	}
	if !jarBind {
		globalClient.Jar = globalCookieJar
		jarBind = true
	}

	temp := EncodeDatas(datas)
	postBuffer := bytes.NewReader(temp)

	request, err = http.NewRequest("POST", _url, postBuffer)
	if err != nil {
		fmt.Println("requset error...")
		return "", ErrorRequest
	}
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Add("Accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Add("Accept-Encoding", "gzip, deflate")

	response, err = globalClient.Do(request)

	if err != nil {
		fmt.Println("response error...")
		return "", ErrorResponse
	}

	defer response.Body.Close()
	var bodyBytes []byte
	bodyBytes, _ = ioutil.ReadAll(response.Body)

	var bodyString string
	bodyString = fmt.Sprintf("%s", bodyBytes)
	Url = request.URL
	//enc := mahonia.NewDecoder("GB18030")
	return bodyString, nil
}
