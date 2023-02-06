package main

import (
	"context"
	"go-grpc-demo/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8002", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageSenderClient(conn)

	oneThreadTest(client)

	fiveThreadsTest(client)
}

func oneThreadTest(client pb.MessageSenderClient) {
	startTime := time.Now()
	for i := 0; i < 100000; i++ {
		resp, err := client.Send(context.Background(), &pb.MessageRequest{
			FirstNum:  32,
			SecondNum: 46,
		})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Println("receive message:", resp.GetResult())
	}
	endTime := time.Now()
	log.Println("oneThreadTest cost time:", startTime.Sub(endTime).Nanoseconds())
}

func fiveThreadsTest(client pb.MessageSenderClient) {
	startTime := time.Now()
	wg := sync.WaitGroup{}
	for core := 0; core < 5; core++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 2000; i++ {
				resp, err := client.Send(context.Background(), &pb.MessageRequest{
					FirstNum:  32,
					SecondNum: 46,
				})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				log.Println("receive message:", resp.GetResult())
			}
		}()
	}
	wg.Wait()
	endTime := time.Now()
	log.Println("fiveThreadsTest cost time:", startTime.Sub(endTime).Nanoseconds())
}
