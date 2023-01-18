package main

import (
	"flag"
	"fmt"

	"github.com/done-devs/proto-go/src/provider"
)

var (
	port = flag.Int("port", 3001, "The server port")
)

func main() {
	fmt.Println("Protobuf files for Go")
	provider.ProviderClient
}