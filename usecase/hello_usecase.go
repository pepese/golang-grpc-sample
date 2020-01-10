package usecase

import (
	"context"
	"log"

	pb "github.com/pepese/golang-grpc-sample/proto/dest/helloworld"
	"google.golang.org/grpc/metadata"
)

type helloUsecase struct{}

func (uc *helloUsecase) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.Name)
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		log.Println(md)
	} else {
		log.Println("Can't get metadata !!")
	}
	return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

func NewHelloUsecase() *helloUsecase {
	return &helloUsecase{}
}
