这是我的公共库

1.XorEnc 异或加解密使用方法：

方法 XorEncodeStr 异或加密，参数msg,key
方法 XorEncodeStr 异或解密，参数msg,key


package main

import "fmt"
import "github.com/xuchengzhi/Library/Encryption"

func main() {
    msg := XorEnc.XorEncodeStr("123456", "abc123testdfdf")
    fmt.Println(msg)
    fmt.Println(XorEnc.XorDecodeStr(msg, "abc123testdfdf"))
}

2.WebSocket 测试

server 启动WebSocket服务，参数 port 

package main

import "github.com/xuchengzhi/Library/WebSocket"

func main() {
    server.Act("8998")
}

3.二维码生成
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

str, name, colors, bolors, x, y 分别为：二维码内容、名称、前景色、背景色 ，宽高

4.获取时间戳和格式化时间
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
