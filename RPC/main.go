package main

import (
	pb "RPC/proto"
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	//gRPC服务地址
	Address = "127.0.0.1:50052"
)

//定义一个helloServer并实现约定的接口
type helloService struct{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "hello" + in.Name + "."
	return resp, nil
}

var HelloServer = helloService{}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		fmt.Printf("failed to listen:%v", err)
	}
	//实现gRPC Server
	s := grpc.NewServer()
	//注册helloServer为客户端提供服务
	pb.RegisterHelloServer(s, HelloServer) //内部调用了s.RegisterServer()
	fmt.Println("Listen on" + Address)

	s.Serve(listen)

}
