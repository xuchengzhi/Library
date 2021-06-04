package UrlRun

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/xuchengzhi/Library/Time"
	"github.com/xuchengzhi/Library/log"
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
	Url    string
	Params map[string]interface{}
}
type ResJson struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Info string `json:"info"`
}

type RunInfo struct {
	ErrorRate string `json:"错误率"`
	RunAvg    string `json:"平均时间"`
	Errors    int    `json:"错误数"`
	TimeList  string `json:"响应集合"`
	RunTime   string `json:"运行时间"`
	StartTime string `json:"开始时间"`
	MinTime   string `json:"最短响应"`
	MaxTime   string `json:"最长响应"`
	EndTime   string `json:"结束时间"`
	QPS       string `json:"QPS"`
}

type RunInfoTmp struct {
	ErrorRate string `json:"errorRate"`
	RunAvg    string `json:"runRate"`
	Errors    int    `json:"errors"`
	TimeList  string `json:"timeList"`
	RunTime   string `json:"runtime"`
	StartTime string `json:"startTime"`
	MinTime   string `json:"minTime"`
	MaxTime   string `json:"maxTime"`
	EndTime   string `json:"endTime"`
}

type ApiResInfo struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ApiJson struct {
	Status int         `json:"code"`
	Msg    interface{} `json:"msg"`
	Data   interface{} `json:"info"`
	Time   interface{} `json:"responsetime"`
	Dates  interface{} `json:"rundates"`
}

var (
	ProxyUrl                  = "http://192.168.248.249:8001"
	Read_Res, IsProxy, IsJson bool
	timeout                   time.Duration
	timeList                  = []time.Duration{}
	Errors                    int
	timecount                 time.Duration
)

func init() {

	timeout = time.Duration(10000 * time.Millisecond)
}

func Post(p Par, ch chan ApiJson, IsProxy bool) {
	urls := p.Url
	// fmt.Println(urls)
	par := p.Params
	var clusterinfo = url.Values{}
	for key, val := range par {
		clusterinfo.Add(key, fmt.Sprintf("%v", val))
	}

	data := clusterinfo.Encode()
	req, _ := http.NewRequest("POST", urls, strings.NewReader(data))
	ctx, cancel := context.WithTimeout(context.Background(), timeout) //设置超时时间
	// log.Info.Println(cancel)
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", "goTest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if IsJson {
		jsonstr, _ := json.Marshal(par)
		req, _ = http.NewRequest("POST", urls, bytes.NewBuffer(jsonstr))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "goTest")
	}
	proxy := func(*http.Request) (*url.URL, error) {
		// log.Info.Println("使用代理")
		return url.Parse(ProxyUrl)
	}
	transport := &http.Transport{}
	if IsProxy {
		log.Info.Println("使用代理")
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
	timecount += duration
	timeList = append(timeList, duration)
	runtime := fmt.Sprintf("%.03f S", duration.Seconds())

	// fmt.Println(duration)
	if err != nil {
		log.Info.Println("time out  Error", err)
		cancel()
		ch <- ApiJson{2, "time out", "error", runtime, GetTime.TF()}
	} else {

		defer resp.Body.Close()

		body, errs := ioutil.ReadAll(resp.Body)
		log.Info.Println(len(body))
		if errs != nil {
			fmt.Println("errors", errs)
		}

		go Response(body, resp.StatusCode)
		if resp.StatusCode != 200 {
			ch <- ApiJson{1, fmt.Sprintf("api error %v", resp.StatusCode), string(body), runtime, GetTime.TF()}
		} else {
			ch <- ApiJson{0, "success", string(body), runtime, GetTime.TF()}
		}

	}
}

func Response(msg []byte, StatusCode int) {
	var s ApiResInfo

	if Read_Res {
		log.Info.Println("response", string(msg))
		log.Info.Println("打印返回结果")
	} else {
		err := json.Unmarshal(msg, &s)
		if err != nil {
			log.Info.Println("err", err)
		} else {
			if s.Code != 0 || StatusCode != 200 {
				Errors += 1
				log.Info.Println(s)
			}
		}
	}

}

func Get(p Par, ch chan ApiJson, IsProxy bool) {
	urls := p.Url
	// fmt.Println(urls)
	par := p.Params
	req, err := http.NewRequest("GET", urls, nil)

	if err != nil {
		log.Info.Println(err)
		ch <- ApiJson{2, err, "error", "0", GetTime.TF()}
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout) //设置超时时间
	log.Info.Println(cancel)
	req = req.WithContext(ctx)
	req.Header.Set("User-Agent", "goTest")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	q := req.URL.Query()
	for key, val := range par {
		q.Add(key, fmt.Sprintf("%v", val))
	}
	req.URL.RawQuery = q.Encode()
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(ProxyUrl)
	}
	transport := &http.Transport{}
	if IsProxy {
		log.Info.Println("使用代理")
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
	// "0" := t2.Sub(t1)
	var duration time.Duration = t2.Sub(t1)
	runtime := fmt.Sprintf("%.03f", duration.Seconds())
	log.Info.Println(fmt.Sprintf("Wait  [%v] Milliseconds [%d] Seconds [%.3f] ", duration, duration.Nanoseconds()/1e6, duration.Seconds()))
	log.Info.Println(runtime)
	if err != nil {
		log.Info.Println("err", err)
		ch <- ApiJson{2, "request canceled or time out", "error", runtime, GetTime.TF()}
	} else {

		defer resp.Body.Close()
		body, errs := ioutil.ReadAll(resp.Body)
		if errs != nil {
			fmt.Println(errs)
		}
		if resp.StatusCode != 200 {
			ch <- ApiJson{1, fmt.Sprintf("api error %v", resp.StatusCode), string(body), runtime, GetTime.TF()}
		} else {
			ch <- ApiJson{0, "success", string(body), runtime, GetTime.TF()}
		}

	}
}

//封装异步执行请求
func Action(urls, method string, p map[string]interface{}, Run_sync *sync.WaitGroup, IsProxy bool) string {
	var pars Par
	// pars := &Par{urls, p}
	pars.Url = urls
	pars.Params = p
	outputs := make(chan ApiJson)
	var status ApiJson

	if strings.EqualFold(method, "post") {
		go Post(pars, outputs, IsProxy)
	} else {
		go Get(pars, outputs, IsProxy)
	}

	defer Run_sync.Done()
	status = <-outputs

	JsonStr, _ := json.Marshal(status)
	return string(JsonStr)

}
func Max(list []time.Duration) time.Duration {
	if len(list) > 0 {
		for i := 0; i < len(list)-1; i++ {
			for j := range list {
				if j == len(list)-1 {
					break
				}
				var tmp = list[j]
				if list[j] < list[j+1] {
					list[j] = list[j+1]
					list[j+1] = tmp
				}
			}

		}
		return list[0]
	} else if len(list) == 1 {
		return list[0]
	}
	return 0
}

func Min(list []time.Duration) time.Duration {
	if len(list) > 0 {
		for i := 0; i < len(list)-1; i++ {
			for j := range list {
				if j == len(list)-1 {
					break
				}
				var tmp = list[j+1]
				if list[j] > list[j+1] {
					list[j+1] = list[j]
					list[j] = tmp
				}
			}

		}
		return list[0]
	} else if len(list) == 1 {
		return list[0]
	}
	return 0
}

// 并发执行
func PressureRun(num int, url, method string, params map[string]interface{}, IsProxy bool) interface{} {
	var res string
	time1 := time.Now()
	var Run_sync sync.WaitGroup
	for i := 0; i < num; i++ {
		Run_sync.Add(1)
		go func() {
			Action(url, method, params, &Run_sync, IsProxy)
		}()
		time.Sleep(1)
	}
	Run_sync.Wait()
	time2 := time.Now()
	time.Sleep(2 * time.Second)

	var runinfo RunInfo

	error_rate := float64(Errors) / float64(num) * 100

	run_rate := int64(timecount) / int64(num)

	runinfo.ErrorRate = fmt.Sprintf("%v%%", error_rate)
	runinfo.RunAvg = fmt.Sprintf("%v毫秒", run_rate)
	if run_rate > 1000000 {
		runinfo.RunAvg = fmt.Sprintf("%.4f毫秒", float64(run_rate)/1000000)
	}
	runinfo.StartTime = time1.Format("2006-01-02 15:04:05 06")
	runinfo.EndTime = time2.Format("2006-01-02 15:04:05 06")
	runinfo.Errors = Errors
	runinfo.TimeList = fmt.Sprintf("%v", timeList)
	runinfo.RunTime = fmt.Sprintf("%v", time2.Sub(time1))
	runinfo.MinTime = fmt.Sprintf("%v", Min(timeList))
	runinfo.MaxTime = fmt.Sprintf("%v", Max(timeList))
	runinfo.QPS = fmt.Sprintf("%v", float64(num)/(float64(run_rate)/1000))

	if Errors != 0 {
		res = fmt.Sprintf("错误数：%v，错误率：%.2f%%,运行时间：%v,响应时间 %v", Errors, error_rate, time2.Sub(time1), timeList)

	} else {
		res = fmt.Sprintf("错误数：%v，错误率：%.2f%%,运行时间：%v,响应时间 %v", Errors, error_rate, time2.Sub(time1), timeList)
	}
	log.Info.Println(timeList)
	timeList = timeList[0:0]
	log.Info.Println(res)
	Errors = 0
	timecount = 0.0
	return runinfo
}
