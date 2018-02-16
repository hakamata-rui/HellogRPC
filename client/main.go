package main

import (
	"fmt"
	"github.com/hakamata-rui/HellogRPC/pb"
	"github.com/prometheus/common/log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"os"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	defer conn.Close()
	c := pb.NewGreedClient(conn)
	message := &pb.HelloRequest{"world"}
	res, err := c.Hello(context.TODO(), message)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	fmt.Println(res)
}
