这是我的公共库

1.XorEnc 异或加解密使用方法：

方法 XorEncodeStr 异或加密，参数msg,key
方法 XorEncodeStr 异或解密，参数msg,key


package main

import "fmt"
import "github.com/xuchengzhi/Library"

func main() {
    msg := XorEnc.XorEncodeStr("123456", "abc123testdfdf")
    fmt.Println(msg)
    fmt.Println(XorEnc.XorDecodeStr(msg, "abc123testdfdf"))
}