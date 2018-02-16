package main

import (
	"log"
	"net"

	"github.com/hakamata-rui/HellogRPC/pb"
	"github.com/hakamata-rui/HellogRPC/service"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	s := &service.GreedService{}
	// 実行したい実処理をseverに登録する
	pb.RegisterGreedServer(server, s)
	server.Serve(listenPort)
}
