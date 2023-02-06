package main

import (
	"context"
	"go-grpc-demo/grpc/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageSenderClient(conn)
	resp, err := client.Send(context.Background(), &pb.MessageRequest{SaySomething: "hello world!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println("receive message:", resp.GetResponseSomething())
}
