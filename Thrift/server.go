package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	// "github.com/xuchengzhi/apimonitor/thriftserver/gen-go/demo/rpc"
	"github.com/xuchengzhi/apimonitor/thriftserver/gen-go/testone/rpc"
	"os"
)

const (
	NetworkAddr = "127.0.0.1:19090"
)

type RpcServiceImpl struct {
}

func (this *RpcServiceImpl) FunCall(callTime int64, funCode string, paramMap map[string]string) (r []string, err error) {
	fmt.Println("-->FunCall:", callTime, funCode, paramMap)

	for k, v := range paramMap {
		r = append(r, k+v)
	}
	return
}

func (this *RpcServiceImpl) TestOne(msg, ips string) (r string, err error) {

	r = fmt.Sprintf("server:%v,ip:%v", msg, ips)
	return
}

func main() {
	// transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transportFactory := thrift.NewTBufferedTransportFactory(10000000)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	//protocolFactory := thrift.NewTCompactProtocolFactory()

	serverTransport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &RpcServiceImpl{}
	processor := rpc.NewRpcServiceProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", NetworkAddr)
	server.Serve()
}
