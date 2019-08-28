package main

// package Thrift

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	// "github.com/xuchengzhi/apimonitor/thriftserver/gen-go/demo/rpc"
	"github.com/xuchengzhi/apimonitor/thriftserver/gen-go/testone/rpc"
	"net"
	"os"
	"time"
)

func Run(msg, ips string) string {

	// transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "19090"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := rpc.NewRpcServiceClientFactory(useTransport, protocolFactory)
	err = transport.Open()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:19090", " ", err)
		// os.Exit(1)
		// return "服务未启动"
	}
	defer transport.Close()

	// for i := 0; i < 1000; i++ {
	// 	paramMap := make(map[string]string)
	// 	paramMap["name"] = "qinerg"
	// 	paramMap["passwd"] = "123456"
	// 	r1, e1 := client.FunCall(currentTimeMillis(), "login", paramMap)
	// 	fmt.Println(i, "Call->", r1, e1)
	// }
	r1, e1 := client.TestOne(msg, ips)
	fmt.Println(r1, e1)
	if e1 == nil {
		return r1
	} else {
		fmt.Println(e1.Error())
		return "服务未启动"
	}
}

func main() {
	startTime := currentTimeMillis()
	msg := Run("ceshi", "djdokdso")
	fmt.Println(msg)
	endTime := currentTimeMillis()
	fmt.Println("Program exit. time->", endTime, startTime, (endTime - startTime))
}

// 转换成毫秒
func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
