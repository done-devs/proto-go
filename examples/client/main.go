package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/done-devs/proto-go/provider"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:3001", "the address to connect to")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := provider.NewProviderClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetId(ctx, &provider.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Server Id: %s", r.Value)
}
