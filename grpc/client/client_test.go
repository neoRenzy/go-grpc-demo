package main

import (
	"context"
	"go-grpc-demo/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func BenchMarkOneThread(b *testing.B) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageSenderClient(conn)

	b.ResetTimer()
	for i := 0; i < MAX_WORK_NUM; i++ {
		_, err := client.Send(context.Background(), &pb.MessageRequest{
			FirstNum:  FirstNum,
			SecondNum: SecondNum,
		})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}
}
