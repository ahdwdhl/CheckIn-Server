package logbuf

import (
	"fmt"
	"log"
	"os"
)

func GetLogFileBuffer(name string) *os.File{
	logFile, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("open file error")
		os.Exit(-1)
	}
	return logFile
}
func GetLogger(name string) (*log.Logger, *os.File){
	logFile := GetLogFileBuffer(name)
	logger := log.New(logFile, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)
	return logger, logFile
}