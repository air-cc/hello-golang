package impl

import (
	"context"
	"log"

	"iaircc.com/go/demo/hellogrpc/helloworld"
)

type GreeterService struct {
	helloworld.UnimplementedGreeterServer
}

// NewGreeterService new a greeter service.
func NewGreeterService() *GreeterService {
	return &GreeterService{}
}

func (s *GreeterService) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}
