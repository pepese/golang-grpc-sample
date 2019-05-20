package main

import (
	"log"
	"net"

	pb "github.com/pepese/golang-grpc-sample/proto/dest/helloworld"
	"github.com/pepese/golang-grpc-sample/usecase"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, usecase.NewHelloUsecase())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
