package GetTime

// package main

import "time"
import "fmt"
import "runtime"
import "reflect"

//获取方法运行时间timecheck
func TC(f func()) time.Duration {
	fname := fmt.Sprintf("func name is %s", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
	time1 := time.Now()
	acttime := time.Now().String()
	f()
	time2 := time.Now()
	endtime := time.Now().String()
	runtime := fmt.Sprintf("start:%v,end:%v", acttime[:24], endtime[:24])
	fmt.Println(runtime)
	fmt.Println(fname)

	return time2.Sub(time1)
}

//获取格式化时间TimeFormat
func TF() string {
	times := time.Now().Format("2006-01-02 15:04:05")
	return times
}

//获取时间戳
func TS() string {
	t := fmt.Sprintf("%v", time.Now().UnixNano()/1e6)
	return t
}

// func main() {
// 	t := TC(test)
// 	fmt.Println(t)
// }
