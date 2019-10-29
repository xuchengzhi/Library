package UrlRun

// package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/xuchengzhi/Library/Time"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"
	"time"
)

//定义结构体http请求url和参数
type Par struct {
	Url    string
	Params map[string]string
}

type ApiJson struct {
	Status int         `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"info"`
	Time   interface{} `json:"responsetime"`
	Dates  interface{} `json:"rundates"`
}

var timeout time.Duration

var is_res, is_proxy bool

func init() {

	timeout = time.Duration(10 * time.Millisecond)
	is_res = true
	is_proxy = true
}

func Post(p Par, ch chan ApiJson, is_proxy bool) {

	urls := p.Url
	// fmt.Println(urls)
	par := p.Params
	var clusterinfo = url.Values{}
	for key, val := range par {
		clusterinfo.Add(key, string(val))
	}

	data := clusterinfo.Encode()
	req, _ := http.NewRequest("POST", urls, strings.NewReader(data))
	ctx, _ := context.WithTimeout(context.Background(), timeout) //设置超时时间
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", "goTest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// proxy := func(_ *http.Request) (*url.URL, error) {
	// 	return url.Parse("http://127.0.0.1:8888")
	// }
	// transport := &http.Transport{}
	proxy := func(*http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8888")
	}

	log.Println(reflect.TypeOf(proxy))

	transport := &http.Transport{}
	if is_proxy {
		transport = &http.Transport{Proxy: proxy}
	}

	client := &http.Client{
		// Timeout:   timeout,
		Transport: transport,
	}
	t1 := time.Now()
	resp, err := client.Do(req)
	// resp, err := transport.RoundTrip(req)
	t2 := time.Now()
	// runtime := (t2.Sub(t1))
	var duration time.Duration = t2.Sub(t1)

	runtime := fmt.Sprintf("%.03f S", duration.Seconds())

	// fmt.Println(duration)
	if err != nil {
		ch <- ApiJson{2, "time out", "error", runtime, GetTime.TS()}
	} else {

		defer resp.Body.Close()
		body, errs := ioutil.ReadAll(resp.Body)
		if errs != nil {
			fmt.Println(errs)
		}
		if resp.StatusCode != 200 {
			ch <- ApiJson{1, "api error", string(body), runtime, GetTime.TS()}
		} else {
			ch <- ApiJson{0, string(body), "success", runtime, GetTime.TS()}
		}

	}
}

func Get(p Par, ch chan ApiJson, is_proxy bool) {

	urls := p.Url
	// fmt.Println(urls)
	par := p.Params
	req, _ := http.NewRequest("GET", urls, nil)
	ctx, _ := context.WithTimeout(context.Background(), timeout) //设置超时时间
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", "goTest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	for key, val := range par {
		q.Add(key, string(val))
	}
	req.URL.RawQuery = q.Encode()
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8888")
	}
	transport := &http.Transport{}
	if is_proxy {
		transport = &http.Transport{Proxy: proxy}
	}

	client := &http.Client{
		// Timeout:   timeout,
		Transport: transport,
	}

	t1 := time.Now()
	resp, err := client.Do(req)
	// resp, err := transport.RoundTrip(req)
	t2 := time.Now()
	// runtime := t2.Sub(t1)
	var duration time.Duration = t2.Sub(t1)
	runtime := fmt.Sprintf("%.03f", duration.Seconds())
	fmt.Printf("Wait  [%v]\nMilliseconds [%d]\nSeconds [%.3f]\n", duration, duration.Nanoseconds()/1e6, duration.Seconds())
	fmt.Println(runtime)
	if err != nil {
		ch <- ApiJson{2, "request canceled or time out", "error", runtime, GetTime.TS()}
	} else {

		defer resp.Body.Close()
		body, errs := ioutil.ReadAll(resp.Body)
		if errs != nil {
			fmt.Println(errs)
		}
		if resp.StatusCode != 200 {
			ch <- ApiJson{1, "api error", string(body), runtime, GetTime.TS()}
		} else {
			ch <- ApiJson{0, string(body), "success", runtime, GetTime.TS()}
		}

	}
}

func Action(urls, method string, p map[string]string, Run_sync *sync.WaitGroup) string {
	var pars Par
	// pars := &Par{urls, p}
	pars.Url = urls
	pars.Params = p
	outputs := make(chan ApiJson)
	var status ApiJson
	if strings.EqualFold(method, "post") {
		go Post(pars, outputs, is_proxy)
	} else {
		go Get(pars, outputs, is_proxy)
	}
	defer Run_sync.Done()
	status = <-outputs
	vvvv, _ := json.Marshal(status)
	fmt.Println(string(vvvv))
	return string(vvvv)

}

type ResJson struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Info string `json:"info"`
}

func PressureRun(num int, url, method string, params map[string]string) {
	errors := 0
	var res string
	time1 := time.Now()
	var Run_sync sync.WaitGroup
	for i := 0; i < num; i++ {
		Run_sync.Add(1)
		go func() {
			tmp := Action(url, method, params, &Run_sync)
			var res ResJson
			// fmt.Println(reflect.TypeOf(tmp))
			if err := json.Unmarshal([]byte(tmp), &res); err == nil {
				if res.Code == 0 {

				} else {
					errors += 1
				}
			} else {
				errors += 1
			}
		}()

		time.Sleep(1)
	}
	Run_sync.Wait()
	time2 := time.Now()
	if errors != 0 {
		error_rate := float64(errors) / float64(num) * 100
		res = fmt.Sprintf("错误数：%v，错误率：%.2f%%,运行时间：%v", errors, error_rate, time2.Sub(time1))

	} else {
		res = fmt.Sprintf("已全部执行完成")

	}
	if is_res {
		fmt.Println(res)
	}
}

// func main() {
// 	params := make(map[string]string)
// 	params["name"] = "test"
// 	params["age"] = "ten"
// 	// tmp := UrlRun.Action("http://192.168.248.188:8082/v1/StsToken", params, "get")
// 	// var res ResJson
// 	// fmt.Println(reflect.TypeOf(tmp))
// 	// if err := json.Unmarshal([]byte(tmp), &res); err == nil {
// 	// 	if res.Code == 0 {
// 	// 		fmt.Println("yes")
// 	// 	} else {
// 	// 		fmt.Println("no")
// 	// 	}
// 	// } else {
// 	// 	fmt.Println(err)
// 	// }
// }
