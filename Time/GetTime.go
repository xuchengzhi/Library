package GetTime

// package main

import (
	"errors"
	"fmt"
	// "log"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var now = time.Now()

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
	t := fmt.Sprintf("%v", (time.Now().UnixNano() / 1e6))
	return t[:len(t)-3] + "000"
}

func StrToTimestamp(t string) (string, error) {
	// t := time.Now()
	stamp, err := time.ParseInLocation("2006-01-02 15:04:05", t, time.Local)
	if err != nil {
		return "", errors.New("time format err")
	}
	res := fmt.Sprintf("%v", stamp.UnixNano()/1e6)
	return res, nil
}

func TimestampToStr(t string) (string, error) {
	// t := time.Now()

	times, err := strconv.ParseInt(strings.Replace(t, "000", "", -1), 10, 64)
	if err != nil {
		return "", errors.New("type err")
	}

	stamp := time.Unix(times, 0).Format("2006-01-02 15:04:05")
	return stamp, nil
}

func SecBefore(num time.Duration) time.Time {
	// 多少分钟之前
	now = time.Now()
	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(num * m)
	return m1
}

func SecAfter(num time.Duration) time.Time {
	// 多少分钟之后
	now = time.Now()
	m, _ := time.ParseDuration("1m")
	m1 := now.Add(num * m)
	return m1
}

func HourBefore(num time.Duration) time.Time {
	// 多少小时之前
	now = time.Now()
	h, _ := time.ParseDuration("-1h")
	h1 := now.Add(num * h)
	fmt.Println(now, h1)
	return h1
}

func HourAfter(num time.Duration) time.Time {
	// 多少小时之后
	now = time.Now()
	h, _ := time.ParseDuration("1h")
	h1 := now.Add(num * h)
	fmt.Println(h1)
	return h1
}

func DayAfter(num time.Duration) time.Time {
	// 多少天之后
	now = time.Now()
	d, _ := time.ParseDuration("24h")
	d1 := now.Add(d * num)
	fmt.Println(now, d1)
	return d1
}

func DayBefore(num time.Duration) time.Time {
	// 多少天之前
	d, _ := time.ParseDuration("-24h")
	d1 := now.Add(d * num)
	fmt.Println(d1)
	return d1
}

func MonthAfter(num int) time.Time {
	// 多少天之后

	now = time.Now()
	d1 := now.AddDate(0, int(now.Month())-num, 0)
	return d1
}

func MonthBefore(num int) time.Time {
	// 多少天之前
	now = time.Now()
	d1 := now.AddDate(0, int(now.Month())+num, 0)
	return d1
	return d1
}

// func test() {
// 	times := time.Now().Format("2006-01-02 15:04:05")
// 	res, _ := StrToUnix(times)
// 	UnixToStr(res)
// }

// func main() {
// fmt.Println(MonthBefore(13))
// fmt.Println(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
// }
