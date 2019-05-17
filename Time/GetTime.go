package GetFuncTime

// package main

import "time"
import "fmt"
import "runtime"
import "reflect"

//获取方法运行时间timecheck
func TC(f func()) time.Duration {
	time1 := time.Now()
	fname := fmt.Sprintf("func name is %s", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
	fmt.Println(fname)
	f()
	time2 := time.Now()
	return time2.Sub(time1)
}

//获取格式化时间TimeFormat
func TF() string {
	times := time.Now().Format("2006-01-02 15:04:05")
	return times
}

func TS() string {
	t := fmt.Sprintf("%v", time.Now().UnixNano()/1e6)
	return t
}

// func main() {
// 	t := TS()
// 	fmt.Println(t)
// }
