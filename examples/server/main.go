package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/done-devs/proto-go/provider"
	"google.golang.org/grpc"
)

var (
	Port = flag.Int("port", 3001, "The server port")
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

// ProviderServer is used to implement the methods.
type ProviderServer struct {
	Id          string
	Name        string
	Description string
	Icon        string
	provider.UnimplementedProviderServer
}

// Implement the logic to create a task
func (s *ProviderServer) CreateTask(ctx context.Context, task *provider.Task) (*provider.ProviderResponse, error) {
	taskJSON, err := json.MarshalIndent(task, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nClient request content: %s", string(taskJSON))
	return &provider.ProviderResponse{
		Successful: true,
		Message:    "Task created successfully",
		Task:       task,
	}, nil
}

// Implement the rest of the methods here.
