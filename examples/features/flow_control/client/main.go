package main

import (
	"context"
	"flag"
	"io"
	pb "keshuigu/mygrpc/exapmles/features/proto/echo"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr = flag.String("addr", "localhost:50051", "the address to connect to")
var payload = make([]byte, 1) // 间接控制了阻塞前发送信息的数量

type event struct {
	fired int32
	c     chan struct{}
	o     sync.Once
}

func (e *event) Fire() bool {
	ret := false
	e.o.Do(func() {
		atomic.StoreInt32(&e.fired, 1)
		close(e.c)
		ret = true
	})
	return ret
}

func (e *event) Done() <-chan struct{} {
	return e.c
}

func (e *event) HasFired() bool {
	return atomic.LoadInt32(&e.fired) == 1
}

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := c.BidirectionalStreamingEcho(ctx)

	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}
	log.Printf("New stream began.")

	stopSending := event{c: make(chan struct{})}
	sendOne := make(chan struct{})
	go func() {
		i := 0
		for !stopSending.HasFired() {
			i++
			if err := stream.Send(&pb.EchoRequest{Message: string(payload)}); err != nil {
				log.Fatalf("Error sending data: %v", err)
			}
			sendOne <- struct{}{}
		}
		log.Printf("Sent %v messages.", i)
		stream.CloseSend()
	}()

	for !stopSending.HasFired() {
		after := time.NewTimer(time.Second)
		select {
		case <-after.C:
			log.Printf("sending is blocked.")
			stopSending.Fire()
			<-sendOne
		case <-sendOne:
			after.Stop()
		}
	}

	// Next, we wait 2 seconds before reading from the stream, to give the
	// server an opportunity to block while sending its responses.
	time.Sleep(2 * time.Second)

	// Finally, read all the data sent by the server to allow it to unblock.
	for i := 0; true; i++ {
		if _, err := stream.Recv(); err != nil {
			log.Printf("Read %v messages.", i)
			if err == io.EOF {
				log.Printf("Stream ended successfully.")
				return
			}
			log.Fatalf("Error receiving data: %v", err)
		}
	}
}
