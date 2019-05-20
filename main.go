package main

import (
	"log"
	"net"

	pb "github.com/pepese/golang-grpc-sample/proto/dest/helloworld"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
