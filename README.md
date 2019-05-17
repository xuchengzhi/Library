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