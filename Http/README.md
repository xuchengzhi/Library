**Http常用库**

- Post 使用方法：

type Par struct {
    Url    string 请求链接
    Params map[string]string 请求参数
}


参数 Post(p Par, ch chan ApiJson, is_proxy bool)
    p 请求参数及链接 ch 获取返回结果 is_proxy 是否代理，默认不使用代理，启用代理为127.0.0.1:8888

```bash
    package main

    import (
        "github.com/xuchengzhi/Library/Http"
        "log"
    )

    func main() {
        par := make(map[string]string)
        par["p"] = "ceshi"
        var pars UrlRun.Par
        pars.Url = "http://xx.xx.com/mobile/cdfdf"
        pars.Params = par
        log.Println(pars)
        res := make(chan UrlRun.ApiJson)
        UrlRun.Post(pars, res, false)
        log.Println(res)
    }
```

- Get 使用方法：

type Par struct {
    Url    string 请求链接
    Params map[string]string 请求参数
}


参数 Get(p Par, ch chan ApiJson, is_proxy bool)
    p 请求参数及链接 ch 获取返回结果 is_proxy 是否代理，默认不使用代理，启用代理为127.0.0.1:8888

```bash
    package main

    import (
        "github.com/xuchengzhi/Library/Http"
        "log"
    )

    func main() {
        par := make(map[string]string)
        par["p"] = "ceshi"
        var pars UrlRun.Par
        pars.Url = "http://xxx.xx.com/mobile/cdfdf"
        pars.Params = par
        log.Println(pars)
        res := make(chan UrlRun.ApiJson)
        UrlRun.Get(pars, res, false)
        log.Println(res)
    }
```

- PressureRun 压力测试使用方法：

参数： PressureRun(num int, url, method string, params map[string]string,is_proxy bool)
        num 请求数 url 接口地址 method 请求方式 params 请求参数

```bash
package main

import (
    "github.com/xuchengzhi/Library/Http"
)

func main() {
    UrlRun.ProxyUrl = "http://192.168.248.150:8001"
    IsProxy := true
    UrlRun.IsJson = true
    UrlRun.Read_Res = true
    params := make(map[string]interface{})
    params["msg"] = "123456"
    params["age"] = "ten"
    UrlRun.PressureRun(20, "https://www.xxx.com.cn/v1/tools/MD5", "post", params, IsProxy)
}
```