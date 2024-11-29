package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	ecpb "keshuigu/mygrpc/examples/features/proto/echo"
	hwpb "keshuigu/mygrpc/examples/helloworld/helloworld"
)

var port = flag.Int("port", 50051, "the port to serve on")

// hwServer is used to implement helloworld.GreeterServer.
type hwServer struct {
	hwpb.UnimplementedGreeterServer
}
type ecServer struct {
	ecpb.UnimplementedEchoServer
}

// SayHello implements helloworld.GreeterServer
func (s *hwServer) SayHello(_ context.Context, in *hwpb.HelloRequest) (*hwpb.HelloReply, error) {
	return &hwpb.HelloReply{Message: "Hello " + in.Name}, nil
}

// UnaryEcho implements echo.EchoServer
func (s *ecServer) UnaryEcho(_ context.Context, req *ecpb.EchoRequest) (*ecpb.EchoResponse, error) {
	return &ecpb.EchoResponse{Message: req.Message}, nil
}
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server listening at %v\n", lis.Addr())

	s := grpc.NewServer()

	hwpb.RegisterGreeterServer(s, &hwServer{})

	ecpb.RegisterEchoServer(s, &ecServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
