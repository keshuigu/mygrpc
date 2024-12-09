package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	pb "keshuigu/mygrpc/examples/features/proto/echo"

	"google.golang.org/grpc"
)

var (
	addrs = []string{"localhost:50050", "localhost:50051"}
)

type echoServer struct {
	pb.UnimplementedEchoServer
	addr string
}

func (s *echoServer) UnaryEcho(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: fmt.Sprintf("%s (from %s)", req.Message, s.addr)}, nil
}

func main() {
	var wg sync.WaitGroup
	for _, addr := range addrs {
		lis, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()
		pb.RegisterEchoServer(s, &echoServer{addr: addr})
		log.Printf("serving on %s\n", addr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
		}()
	}
	wg.Wait()
}
