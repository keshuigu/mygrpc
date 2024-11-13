package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	ecpb "keshuigu/mygrpc/exapmles/features/proto/echo"
	hwpb "keshuigu/mygrpc/exapmles/helloworld/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	// flag包用于解析命令行参数
	port = flag.Int("port", 50051, "The server port")
)

type hwServer struct {
	// proto文件结构体
	// 嵌入此处代表帮server结构体实现默认方法
	// 意味着如果没有实现SayHello
	// 嵌入的方法会提升到外层
	hwpb.UnimplementedGreeterServer
}

func (s *hwServer) SayHello(_ context.Context, in *hwpb.HelloRequest) (*hwpb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &hwpb.HelloReply{Message: "Hello" + in.GetName()}, nil
}

func (s *hwServer) SayHelloAgain(_ context.Context, in *hwpb.HelloRequest) (*hwpb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &hwpb.HelloReply{Message: "Hello" + in.GetName()}, nil
}

type ecServer struct {
	ecpb.UnimplementedEchoServer
}

func (s *ecServer) UnaryEcho(_ context.Context, in *ecpb.EchoRequest) (*ecpb.EchoResponse, error) {
	return &ecpb.EchoResponse{Message: in.GetMessage()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hwpb.RegisterGreeterServer(s, &hwServer{})
	ecpb.RegisterEchoServer(s, &ecServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
