package usecase

import (
	"context"
	"log"

	pb "github.com/pepese/golang-grpc-sample/proto/dest/helloworld/v2"
	"google.golang.org/grpc/metadata"
)

type helloUsecaseV2 struct{}

func (uc *helloUsecaseV2) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.Name)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Println(md)
	} else {
		log.Println("Can't get metadata !!")
	}
	return &pb.HelloResponse{Message: "Hello " + in.Name + " v2"}, nil
}

func NewHelloUsecaseV2() *helloUsecaseV2 {
	return &helloUsecaseV2{}
}
