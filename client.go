package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "example.com/myapp/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

// make sure to edit the addr url below, i had my code deployed to Azure Container Apps, for local testing you can use localhost
var (
	addr = flag.String("addr", "websocket-grpc-app-witty.wittyhill-7980bc95.azurecontainerapps-test.io:9091", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response. The 1 second below means channel will stay open for 1 second then time out, increasing that number increases the amout of times we can comm over channel
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
