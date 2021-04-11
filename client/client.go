package main

/**
created lhs
*/
import (
	"context"
	"google.golang.org/grpc"
	"grcp/helloworld"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)

	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)
	name := "Test"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
	if err != nil {
		log.Fatal("could not greet: %v", err)
	}
	log.Println("Greeting:%s", r.Message)
}
