package UrlRun

// package main

import (
	"fmt"
	// "github.com/xuchengzhi/Library/Time"
	// "context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	// "reflect"
	"strings"
	"sync"
	"time"
)

//定义结构体http请求url和参数
type Par struct {
	url    string
	params map[string]string
}

type ApiJson struct {
	Status int         `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"info"`
	Time   interface{} `json:"time"`
}

var timeout time.Duration

var is_res, is_proxy bool

func init() {

	timeout = time.Duration(500 * time.Millisecond)
	is_res = true
	is_proxy = false
}

func Post(p *Par, ch chan ApiJson) {

	urls := p.url
	// fmt.Println(urls)
	par := p.params
	var clusterinfo = url.Values{}
	for key, val := range par {
		clusterinfo.Add(key, string(val))
	}

	data := clusterinfo.Encode()
	req, _ := http.NewRequest("POST", urls, strings.NewReader(data))
	req.Header.Set("User-Agent", "goTest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// proxy := func(_ *http.Request) (*url.URL, error) {
	// 	return url.Parse("http://127.0.0.1:8888")
	// }
	// transport := &http.Transport{}
	// proxy := func(_ *http.Request) (*url.URL, error) {
	// 	return url.Parse("http://127.0.0.1:8888")
	// }

	// transport := &http.Transport{Proxy: proxy}
	client := &http.Client{
		Timeout: timeout,
		// Transport: transport,
	}
	t1 := time.Now()
	resp, err := client.Do(req)
	// resp, err := transport.RoundTrip(req)
	t2 := time.Now()
	// runtime := (t2.Sub(t1))
	var duration time.Duration = t2.Sub(t1)

	runtime := fmt.Sprintf("%.03f", duration.Seconds())
	// fmt.Printf("Wait  [%v]\nMilliseconds [%d]\nSeconds [%.3f]\n", duration, duration.Nanoseconds()/1e6, duration.Seconds())
	fmt.Println(duration)
	if err != nil {
		ch <- ApiJson{2, "request canceled or time out", "error", runtime}
	} else {

		defer resp.Body.Close()
		body, errs := ioutil.ReadAll(resp.Body)
		if errs != nil {
			fmt.Println(errs)
		}
		if resp.StatusCode != 200 {
			ch <- ApiJson{1, "api error", string(body), runtime}
		} else {
			ch <- ApiJson{0, string(body), "success", runtime}
		}

	}
}

func Get(p *Par, ch chan ApiJson) {

	urls := p.url
	// fmt.Println(urls)
	par := p.params
	req, _ := http.NewRequest("GET", urls, nil)
	req.Header.Set("User-Agent", "goTest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	for key, val := range par {
		q.Add(key, string(val))
	}
	req.URL.RawQuery = q.Encode()
	// proxy := func(_ *http.Request) (*url.URL, error) {
	// 	return url.Parse("http://127.0.0.1:8888")
	// }
	transport := &http.Transport{}

	client := &http.Client{
		Timeout:   timeout,
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
		ch <- ApiJson{2, "request canceled or time out", "error", runtime}
	} else {

		defer resp.Body.Close()
		body, errs := ioutil.ReadAll(resp.Body)
		if errs != nil {
			fmt.Println(errs)
		}
		if resp.StatusCode != 200 {
			ch <- ApiJson{1, "api error", string(body), runtime}
		} else {
			ch <- ApiJson{0, string(body), "success", runtime}
		}

	}
}

func Action(urls, method string, p map[string]string, Run_sync *sync.WaitGroup) string {

	pars := &Par{urls, p}
	outputs := make(chan ApiJson)
	var status ApiJson
	if strings.EqualFold(method, "post") {
		go Post(pars, outputs)
	} else {
		go Get(pars, outputs)
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
	if errors != 0 {
		error_rate := float64(errors) / float64(num) * 100
		res = fmt.Sprintf("错误数：%v，错误率：%.2f%%", errors, error_rate)

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
