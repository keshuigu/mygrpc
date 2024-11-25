package main

import (
	"context"
	"flag"
	"fmt"
	pb "keshuigu/mygrpc/examples/helloworld/helloworld"
	"log"
	"net"

	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var port = flag.Int("port", 50051, "port number")

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if in.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "request missing required field: Name")
	}
	return &pb.HelloReply{Message: "Hello" + in.Name}, nil
}

func main() {
	flag.Parse()

	address := fmt.Sprintf(":%v", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
