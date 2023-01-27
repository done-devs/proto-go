package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"time"

	"github.com/done-devs/proto-go/provider"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
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
	example_note := "Example notes"
	r, err := c.CreateTask(ctx, &provider.Task{
		Title:                "Example Title",
		Favorite:             false,
		Today:                false,
		Status:               provider.Status_NOT_STARTED,
		Priority:             provider.Priority_HIGH,
		SubTasks:             make([]*provider.SubTask, 0),
		Tags:                 make([]string, 0),
		Notes:                &example_note,
		CompletionDate:       nil,
		DeletionDate:         nil,
		DueDate:              nil,
		ReminderDate:         nil,
		Recurrence:           nil,
		CreatedDateTime:      timestamppb.Now(),
		LastModifiedDateTime: timestamppb.Now(),
	})
	if err != nil {
		log.Fatalf("Could not process the request: %v", err)
	}
	taskJSON, err := json.MarshalIndent(r.Task, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("\nServer response: %s", string(taskJSON))
}
