package main

import (
	"flag"
	"fmt"
	"os"
)

// 实际中应该用更好的变量名
var (
	h    bool
	test string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.String("test", "Default", "获取数据")

	// 改变默认的 Usage，flag包中的Usage 其实是一个函数类型。这里是覆盖默认函数实现，具体见后面Usage部分的分析
	flag.Usage = usage
}

//打印值的函数
func Formats(f *flag.Flag) string {
	res := ""
	if f != nil {
		res = fmt.Sprintf("%v", f.Value)
	} else {
		fmt.Println(nil)
	}
	return res
}

func main() {

	flag.Parse()
	if h {
		flag.Usage()
	}
	msg := Formats(flag.Lookup("test"))
	fmt.Println(msg)
}

func usage() {
	fmt.Fprintf(os.Stderr, `App version: Test/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`)
	flag.PrintDefaults()
}
