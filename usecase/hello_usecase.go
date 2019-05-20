package usecase

import (
	"context"
	"log"

	pb "github.com/pepese/golang-grpc-sample/proto/dest/helloworld"
)

type helloUsecase struct{}

func (uc *helloUsecase) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}
