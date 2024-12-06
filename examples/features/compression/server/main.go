package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	// 安装 gzip 编码会将其注册为可用的压缩器。
	// 如果客户端支持 gzip，gRPC 将自动协商并使用 gzip。
	// "_" 导入为了执行init()
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"

	pb "keshuigu/mygrpc/examples/features/proto/echo"
)

var port = flag.Int("port", 50051, "the port to serve on")

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) UnaryEcho(_ context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	fmt.Printf("UnaryEcho called with message %q\n", in.GetMessage())
	return &pb.EchoResponse{Message: in.Message}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server listening at %v\n", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	s.Serve(lis)
}
