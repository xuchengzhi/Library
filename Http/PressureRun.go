package main

import "encoding/json"
import "github.com/xuchengzhi/Library/Http"
import "time"
import "fmt"

type ResJson struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Info string `json:"info"`
}

var is_res, is_proxy bool

func init() {
	is_res = true
	is_proxy = false
}

func Action(num int, url, method string, params map[string]string) {
	errors := 0
	var res string
	for i := 0; i < num; i++ {

		go func() {
			tmp := UrlRun.Action(url, params, method)
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

func main() {
	params := make(map[string]string)
	params["name"] = "test"
	params["age"] = "ten"
	Action(1000, "http://localhost:8082/v1/StsToken", "post", params)
}
