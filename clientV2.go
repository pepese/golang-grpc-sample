package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/pepese/golang-grpc-sample/proto/dest/helloworld/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	ctx = metadata.AppendToOutgoingContext(ctx, "HOGE", "hoge")
	client := pb.NewHelloServiceClient(conn)
	message := &pb.HelloRequest{Name: "hoge"}
	res, err := client.SayHello(ctx, message)

	fmt.Printf("result:%#v \n", res)
	fmt.Printf("error::%#v \n", err)
}
