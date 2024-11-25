package main

import (
	"flag"
	"fmt"
	"io"
	pb "keshuigu/mygrpc/examples/features/proto/echo"
	"log"
	"net"
	"sync/atomic"
	"time"

	"sync"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "the port to serve on")

var payload = string(make([]byte, 8*1024))

type server struct {
	pb.UnimplementedEchoServer
}

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

func (s *server) BidirectionalStreamingEcho(stream pb.Echo_BidirectionalStreamingEchoServer) error {
	log.Printf("New stream began.")
	// 不进行read以堵塞client
	time.Sleep(2 * time.Second)

	for i := 0; true; i++ {
		if _, err := stream.Recv(); err != nil {
			log.Printf("Read %v messages.", i) // 只输出最后一条消息
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
	}

	// send
	stopSending := event{c: make(chan struct{})}
	sentOne := make(chan struct{})
	go func() {
		for !stopSending.HasFired() {
			after := time.NewTimer(time.Second)
			select {
			case <-sentOne:
				after.Stop()
			case <-after.C:
				log.Printf("Sending is blocked.")
				stopSending.Fire()
				<-sentOne
			}
		}
	}()

	i := 0
	for !stopSending.HasFired() {
		i++
		if err := stream.Send(&pb.EchoResponse{Message: payload}); err != nil {
			log.Printf("Error sending data: %v", err)
			return err
		}
		sentOne <- struct{}{}
	}

	log.Printf("Sent %v messages.", i)

	log.Printf("Stream ended successfully.")
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
