package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "keshuigu/mygrpc/examples/features/proto/echo"

	"keshuigu/mygrpc/examples/features/stats_monitoring/statshandler"
)

var port = flag.Int("port", 50051, "the port to serve on")

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) UnaryEcho(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	time.Sleep(2 * time.Second)
	return &pb.EchoResponse{Message: req.Message}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen on port %d: %v", *port, err)
	}
	log.Printf("server listening at %v\n", lis.Addr())

	s := grpc.NewServer(grpc.StatsHandler(statshandler.New()))
	pb.RegisterEchoServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
