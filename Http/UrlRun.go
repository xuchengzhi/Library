package UrlRun

import (
	"fmt"
	// "github.com/xuchengzhi/Library/Time"
	"encoding/json"
	"net/http"
	"net/url"
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
	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse("http://127.0.0.1:8888")
	}
	transport := &http.Transport{Proxy: proxy}
	timeout := time.Duration(200 * time.Millisecond)
	client := &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}

	resp, err := client.Do(req)
	// resp, err := transport.RoundTrip(req)
	if err != nil {
		ch <- ApiJson{1, "time out", "error"}
	}
	if resp != nil {
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			ch <- ApiJson{1, "api error", "error"}
		} else {
			ch <- ApiJson{0, "success", "ok"}
		}

	}

}

func Action(url string, p map[string]string) string {
	pars := &Par{url, p}
	outputs := make(chan ApiJson)
	var status ApiJson
	go Post(pars, outputs)
	status = <-outputs
	vvvv, _ := json.Marshal(status)
	fmt.Println(string(vvvv))
	return string(vvvv)

}

// func main() {
// 	params := make(map[string]string)
// 	params["name"] = "test"
// 	params["age"] = "ten"
// 	pars := &Par{"http://192.168.248.188:8082/v1/StsToken", params}
// 	Response(pars)
// }
