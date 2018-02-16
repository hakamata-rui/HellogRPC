package service

import (
	"golang.org/x/net/context"
	"fmt"
	"github.com/hakamata-rui/HellogRPC/pb"
)

type GreedService struct {
}

func (*GreedService) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: fmt.Sprintf("hello, %s", req.Name)}, nil
}
