package log

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var (
	Trace     *log.Logger // 记录所有日志
	Info      *log.Logger // 重要的信息
	Warning   *log.Logger // 需要注意的信息
	Error     *log.Logger // 非常严重的问题
	WriteFile bool        //是否写入文件

)

func CheckLogFile() {
	_, err := os.Stat("logs")
	if err == nil {
		log.Println("目录已存在")
	}
	if os.IsNotExist(err) {
		// log.Println("目录不存在,创建中")
		c_err := os.MkdirAll("logs", os.ModePerm)
		if c_err != nil {
			log.Println("目录创建失败", c_err)
		}
	}
}

func init() {
	now := time.Now()

	if WriteFile {
		CheckLogFile()
		file, err := os.OpenFile(fmt.Sprintf("logs/SYS_%v.log", now.Format("20060102")),
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln("Failed to open error log file:", err)
		}
		Trace = log.New(io.MultiWriter(file, os.Stderr),
			"TRACE: ",
			log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(io.MultiWriter(file, os.Stderr),
			"Info: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Warning = log.New(io.MultiWriter(file, os.Stderr),
			"WARNING: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Error = log.New(io.MultiWriter(file, os.Stderr),
			"ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		Trace = log.New(ioutil.Discard,
			"TRACE: ",
			log.Ldate|log.Ltime|log.Lshortfile)
		Info = log.New(os.Stdout,
			"Info: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Warning = log.New(os.Stdout,
			"WARNING: ",
			log.Ldate|log.Ltime|log.Lshortfile)

		Error = log.New(os.Stdout,
			"ERROR: ",
			log.Ldate|log.Ltime|log.Lshortfile)
	}
}
