package usecase

import (
	"context"
	"log"

	pb "github.com/pepese/golang-grpc-sample/proto/dest/helloworld/v1"
	"google.golang.org/grpc/metadata"
)

type helloUsecaseV1 struct{}

func (uc *helloUsecaseV1) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.Name)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Println(md)
	} else {
		log.Println("Can't get metadata !!")
	}
	return &pb.HelloResponse{Message: "Hello " + in.Name + " v1"}, nil
}

func NewHelloUsecaseV1() *helloUsecaseV1 {
	return &helloUsecaseV1{}
}
