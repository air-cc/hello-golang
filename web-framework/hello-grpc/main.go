package main

import (
	"log"
	"net"

	grpc "google.golang.org/grpc"
	helloworld "iaircc.com/go/demo/hellogrpc/helloworld"
	helloworldImpl "iaircc.com/go/demo/hellogrpc/helloworld/impl"
)

func main() {
	s := grpc.NewServer()

	// 注册 rpc 服务
	helloworld.RegisterGreeterServer(s, helloworldImpl.NewGreeterService())

	lis, err := net.Listen("tcp", ":8008")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
