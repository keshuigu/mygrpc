package main

import (
	"context"
	"flag"
	"fmt"
	pb "keshuigu/mygrpc/exapmles/route_guide/routeguide"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedRouteGuideServer
}

func (s *server) GetFeature(context.Context, *pb.Point) (*pb.Feature, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFeature not implemented")
}
func (s *server) ListFeatures(*pb.Rectangle, grpc.ServerStreamingServer[pb.Feature]) error {
	return status.Errorf(codes.Unimplemented, "method ListFeatures not implemented")
}
func (s *server) RecordRoute(grpc.ClientStreamingServer[pb.Point, pb.RouteSummary]) error {
	return status.Errorf(codes.Unimplemented, "method RecordRoute not implemented")
}
func (s *server) RouteChat(grpc.BidiStreamingServer[pb.RouteNote, pb.RouteNote]) error {
	return status.Errorf(codes.Unimplemented, "method RouteChat not implemented")
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterRouteGuideServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failer to server: %v", err)
	}
}
