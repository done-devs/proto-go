package main

import (
	"context"
	"flag"
	"log"

	provider "github.com/done-devs/proto-go/provider"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	Port = flag.Int("port", 3001, "The server port")
)

// ProviderServer is used to implement provider.GetId.
type ProviderServer struct {
	Id          string
	Name        string
	Description string
	Icon        string
	provider.UnimplementedProviderServer
}

func (s *ProviderServer) GetId(ctx context.Context, rq *provider.Empty) (*wrappers.StringValue, error) {
	log.Printf("Received request GetId")
	return wrapperspb.String(s.Id), nil
}

// Implement the rest of the methods here.
