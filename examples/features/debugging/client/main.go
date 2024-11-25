package main

import (
	"context"
	"flag"
	"log"
	"net"
	"time"

	pb "keshuigu/mygrpc/examples/helloworld/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

const defaultName = "world"

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	// 这个服务器是用来服务Channelz的
	flag.Parse()
	lis, err := net.Listen("tcp", *addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	s := grpc.NewServer()
	service.RegisterChannelzServiceToServer(s)
	go s.Serve(lis)
	defer s.Stop()
	// manual.NewBuilderWithScheme("whatever")用于创建一个新的手动解析器构建器，并指定一个自定义的方案。
	// 手动解析器允许你手动控制gRPC客户端的解析过程，例如动态更新服务端点。
	/***** Initialize manual resolver and Dial *****/
	r := manual.NewBuilderWithScheme("whatever")
	r.InitialState(resolver.State{Addresses: []resolver.Address{{Addr: ":10001"}, {Addr: ":10002"}, {Addr: ":10003"}}})
	// Set up a connection to the server.
	conn, err := grpc.NewClient(r.Scheme()+":///test.server", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithResolvers(r), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	// conn, err := grpc.Dial(r.Scheme()+":///test.server", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithResolvers(r), grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// 连接服务器并打印结果
	// 代理会影响结果
	for i := 0; i < 100; i++ {
		// 服务器设置的是100~200ms
		ctx, cancel := context.WithTimeout(context.Background(), 150*time.Millisecond)
		defer cancel()
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
		if err != nil {
			log.Printf("could not greet: %v", err)
		} else {
			log.Printf("Greeting: %s", r.Message)
		}
	}

	select {}
}
