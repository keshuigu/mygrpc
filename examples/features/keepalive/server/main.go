package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "keshuigu/mygrpc/examples/features/proto/echo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var port = flag.Int("port", 50052, "port number")

// EnforcementPolicy用于在服务器端设置keepalive强制策略。服务器将关闭与违反此策略的客户端的连接。
var kaep = keepalive.EnforcementPolicy{
	// If a client pings more than once every 5 seconds, terminate the connection
	MinTime: 5 * time.Second, // ping 发送最小间隔
	// 是否允许没有stream时发送ping
	PermitWithoutStream: true, // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle: 30 * time.Second, // 空闲15秒，发送GOAWAY
	MaxConnectionAge:  60 * time.Second, // 最长活跃时间30秒
	// // 达到 MaxConnectionAge 但未处理完后，服务器将允许连接在宽限时间内继续处理未完成的请求。超过宽限时间后，服务器将关闭连接。
	MaxConnectionAgeGrace: 5 * time.Second, // 5秒宽限时间
	Time:                  5 * time.Second, // 无活动下ping包5秒发送一次
	Timeout:               1 * time.Second, // 等待ping ack的超时时间为1秒
}

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) UnaryEcho(_ context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: in.Message}, nil
}

func main() {
	flag.Parse()
	address := fmt.Sprintf(":%v", *port)
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
	pb.RegisterEchoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
