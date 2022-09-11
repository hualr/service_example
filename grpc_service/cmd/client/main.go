package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "service_example/grpc_service/api/message"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("")
		}
	}(conn)
	client := pb.NewHealthServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	health, err := client.CheckHealth(ctx, &pb.MessageRequest{SaySomething: "test"})
	if err != nil {
		log.Fatalf("error")
	}
	log.Printf(health.ResponseSomething)

}
