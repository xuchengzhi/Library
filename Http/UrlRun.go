package UrlRun

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
}

var timeout time.Duration

func init() {

	timeout = time.Duration(500 * time.Millisecond)
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
	transport := &http.Transport{}

	client := &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	resp, err := client.Do(req)
	// resp, err := transport.RoundTrip(req)
	if err != nil {
		ch <- ApiJson{2, "request canceled or time out", "error"}
	} else {

		defer resp.Body.Close()
		body, errs := ioutil.ReadAll(resp.Body)
		if errs != nil {
			fmt.Println(errs)
		}
		if resp.StatusCode != 200 {
			ch <- ApiJson{1, "api error", string(body)}
		} else {
			ch <- ApiJson{0, string(body), "success"}
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

	resp, err := client.Do(req)
	// resp, err := transport.RoundTrip(req)
	if err != nil {
		ch <- ApiJson{2, "request canceled or time out", "error"}
	} else {
		defer resp.Body.Close()
		body, errs := ioutil.ReadAll(resp.Body)
		if errs != nil {
			fmt.Println(errs)
		}
		if resp.StatusCode != 200 {
			ch <- ApiJson{1, "api error", string(body)}
		} else {
			ch <- ApiJson{0, string(body), "success"}
		}

	}
}

func Action(urls string, p map[string]string, method string) string {

	pars := &Par{urls, p}
	outputs := make(chan ApiJson)
	var status ApiJson
	if strings.EqualFold(method, "post") {
		go Post(pars, outputs)
	} else {
		go Get(pars, outputs)
	}

	status = <-outputs
	vvvv, _ := json.Marshal(status)
	// fmt.Println(string(vvvv))
	return string(vvvv)

}

// func main() {
// 	params := make(map[string]string)
// 	params["name"] = "test"
// 	params["age"] = "ten"
// 	tmp := UrlRun.Action("http://192.168.248.188:8082/v1/StsToken", params, "get")
// 	var res ResJson
// 	fmt.Println(reflect.TypeOf(tmp))
// 	if err := json.Unmarshal([]byte(tmp), &res); err == nil {
// 		if res.Code == 0 {
// 			fmt.Println("yes")
// 		} else {
// 			fmt.Println("no")
// 		}
// 	} else {
// 		fmt.Println(err)
// 	}
// }
