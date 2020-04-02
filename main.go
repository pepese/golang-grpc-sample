package main

import (
	"log"
	"net"

	pb "github.com/pepese/golang-grpc-sample/proto/dest/helloworld"
	pbv1 "github.com/pepese/golang-grpc-sample/proto/dest/helloworld/v1"
	pbv2 "github.com/pepese/golang-grpc-sample/proto/dest/helloworld/v2"
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
	pbv1.RegisterHelloServiceServer(s, usecase.NewHelloUsecaseV1())
	pbv2.RegisterHelloServiceServer(s, usecase.NewHelloUsecaseV2())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
