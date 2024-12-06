package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "keshuigu/mygrpc/examples/features/proto/echo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip" // Install the gzip compressor

	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func main() {
	flag.Parse()

	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)

	// 发送压缩的 RPC。如果客户端上的所有 RPC 都应以这种方式发送，请使用 DialOption：
	// grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name))
	const msg = "compress"
	msg1 := ""
	for i := 0; i < 100; i++ {
		msg1 += msg
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()
	res, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: msg1}, grpc.UseCompressor(gzip.Name))
	fmt.Printf("UnaryEcho call returned %q, %v\n", res.GetMessage(), err)
	// data = 803 payload = 42
	if err != nil || res.GetMessage() != msg1 {
		log.Fatalf("Message=%q, err=%v; want Message=%q, err=<nil>", res.GetMessage(), err, msg)
	}
	// data = 803 payload = 803
	res, err = c.UnaryEcho(ctx, &pb.EchoRequest{Message: msg1})
	fmt.Printf("UnaryEcho call returned %q, %v\n", res.GetMessage(), err)
	if err != nil || res.GetMessage() != msg1 {
		log.Fatalf("Message=%q, err=%v; want Message=%q, err=<nil>", res.GetMessage(), err, msg)
	}
}
