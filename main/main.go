package main

import (
	"../logbuf"
	"../websever"
)

func main() {

	//cookie, _:= utils.ReadEasyCookie()
	//fmt.Println(cookie)
	logger, file := logbuf.GetLogger("test.log")
	logger.Println("日志已经开启")

	websever.Start(logger)
	defer file.Close()

}
