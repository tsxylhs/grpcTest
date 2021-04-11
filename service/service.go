package main

/**
created lhs
*/

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grcp/helloworld"
	"log"
	"net"
)

type SayHelloTest struct {
}

func (this *SayHelloTest) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {

	return &helloworld.HelloReply{Message: "Hello :" + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {

	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &SayHelloTest{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
