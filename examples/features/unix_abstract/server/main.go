package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "keshuigu/mygrpc/examples/features/proto/echo"

	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "abstract-unix-socket", "The unix abstract socket address")
)

type ecServer struct {
	pb.UnimplementedEchoServer
	addr string
}

func (s *ecServer) UnaryEcho(_ context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: fmt.Sprintf("%s (from %s)", req.Message, s.addr)}, nil
}

func main() {
	flag.Parse()
	netw := "unix"
	socketAddr := fmt.Sprintf("@%v", *addr)
	lis, err := net.Listen(netw, socketAddr)
	if err != nil {
		log.Fatalf("net.Listen(%q, %q) failed: %v", netw, socketAddr, err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &ecServer{addr: socketAddr})
	log.Printf("serving on %s\n", lis.Addr().String())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
