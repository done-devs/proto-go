package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/done-devs/proto-go/provider"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	provider.RegisterProviderServer(s, &ProviderServer{
		Id:          "TestServiceId",
		Name:        "Test Service",
		Description: "Test Service Description",
		Icon:        "test-service-icon",
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
