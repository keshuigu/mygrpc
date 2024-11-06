package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "keshuigu/mygrpc/exapmles/helloworld/helloworld"

	"google.golang.org/grpc"
)

var (
	// flag包用于解析命令行参数
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	// proto文件结构体
	// 嵌入此处代表帮server结构体实现默认方法
	// 意味着如果没有实现SayHello
	// 嵌入的方法会提升到外层
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello" + in.GetName()}, nil
}

func (s *server) SayHelloAgain(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello" + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
