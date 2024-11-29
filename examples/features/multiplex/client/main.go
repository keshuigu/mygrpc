package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	ecpb "keshuigu/mygrpc/examples/features/proto/echo"
	hwpb "keshuigu/mygrpc/examples/helloworld/helloworld"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")

func callSayHello(c hwpb.GreeterClient, name string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &hwpb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("client.SayHello(_) = _, %v", err)
	}
	fmt.Println("Greeting:", r.Message)
}
func callUnaryEcho(client ecpb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := client.UnaryEcho(ctx, &ecpb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("client.UnaryEcho(_) = _, %v: ", err)
	}
	fmt.Println("UnaryEcho: ", resp.Message)
}
func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	fmt.Println("--- calling helloworld.Greeter/SayHello ---")
	hwc := hwpb.NewGreeterClient(conn)
	callSayHello(hwc, "multiplex")
	fmt.Println()
	fmt.Println("--- calling routeguide.RouteGuide/GetFeature ---")
	// Make a routeguide client with the same ClientConn.
	rgc := ecpb.NewEchoClient(conn)
	callUnaryEcho(rgc, "this is examples/multiplex")
	callUnaryEcho(rgc, "this is examples/multiplex")
}
