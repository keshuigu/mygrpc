package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	ecpb "keshuigu/mygrpc/examples/features/proto/echo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// A dial target of `unix:@abstract-unix-socket` should also work fine for
	// this example because of golang conventions (net.Dial behavior). But we do
	// not recommend this since we explicitly added the `unix-abstract` scheme
	// for cross-language compatibility.
	addr = flag.String("addr", "abstract-unix-socket", "The unix abstract socket address")
)

func callUnaryEcho(c ecpb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UnaryEcho(ctx, &ecpb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r.Message)
}

func makeRPCs(cc *grpc.ClientConn, n int) {
	hwc := ecpb.NewEchoClient(cc)
	for i := 0; i < n; i++ {
		callUnaryEcho(hwc, "this is examples/unix_abstract")
	}
}

func main() {
	flag.Parse()
	sockAddr := fmt.Sprintf("unix-abstract:%v", *addr)
	cc, err := grpc.NewClient(sockAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.NewClient(%q) failed: %v", sockAddr, err)
	}
	defer cc.Close()

	fmt.Printf("--- calling echo.Echo/UnaryEcho to %s\n", sockAddr)
	makeRPCs(cc, 10)
	fmt.Println()
}
