**这是我的公共库**

- XorEnc 异或加解密使用方法：

方法 XorEncodeStr 异或加密，参数msg,key
方法 XorEncodeStr 异或解密，参数msg,key

```bash
package main

import "fmt"
import "github.com/xuchengzhi/Library/Encryption"

func main() {
    msg := XorEnc.XorEncodeStr("123456", "abc123testdfdf")
    fmt.Println(msg)
    fmt.Println(XorEnc.XorDecodeStr(msg, "abc123testdfdf"))
}
```
- WebSocket 测试

server 启动WebSocket服务，参数 port 
```bash
package main

import "github.com/xuchengzhi/Library/WebSocket"

func main() {
    server.Act("8998")
}
```
- 二维码生成
> str, name, colors, bolors, x, y 分别为：二维码内容、名称、前景色、背景色 ，宽高
```bash
package main

import "github.com/xuchengzhi/Library/Qrcode"

func main() {
    x, y := 500, 500
    name := "ceshi"
    str := "http://www.baidu.com"
    bolors := "FF1493"
    colors := "000000"
    Qr.Builds(str, name, colors, bolors, x, y)
}
```
![enter description here](https://github.com/xuchengzhi/Library/blob/master/Qrcode/ceshi.png)

- 获取时间戳和格式化时间
```bash
package main

import "github.com/xuchengzhi/Library/Time"
import "fmt"

func test() {
    fmt.Println("Hello World")
}

func main() {
    times := GetTime.TC(test)
    timestmap := GetTime.TS()
    timesformat := GetTime.TF()
    fmt.Println(times)
    fmt.Println(timestmap)
    fmt.Println(timesformat)
}
```

- 获取apk/ipa 包名版本号
```bash
package main

import "github.com/xuchengzhi/Library/AppInfo"
import "fmt"

func main() {
    Istats, IOS := CheckApp.IOS("E:/code/py/shoujizaozi_Test/pachong/study/appdown/App/IOS/9b497ab9-69b3-4dd1-b97b-58cca6bf339a.ipa")
    Astats, ADR := CheckApp.Adr("./ceshi.apk")
    if Istats {
        fmt.Printf("IOS包名%v\n", IOS.Name)
        fmt.Printf("版本号%v\n", IOS.Version)
    }
    if Astats {
        fmt.Printf("IOS包名%v\n", ADR.Name)
        fmt.Printf("版本号%v\n", ADR.Version)
    }

}

```
```bash
> Output:
IOS包名me.myfont.HandFontMaker
版本号4.7.0
安卓包名com.handwriting.makefont
版本号4.7.0
```

- 获取随机字符串
```bash
package main

import (
    "github.com/xuchengzhi/Library/Random"
    "log"
    )
func main() {
    msg := Randoms.GetRandomInt(10000, 99999) //GetRandomString(18)
    log.Println(msg)
}
```

- 进制转换

进制转换引用 https://blog.csdn.net/flyfreelyit/article/details/80097035

```bash
package main

import (
    "github.com/xuchengzhi/Library/Transfar"
    "log"
    )
func main() {
    log.Println(DecBin(10))
}
```