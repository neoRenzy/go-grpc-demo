package main

import (
	"context"
	"fmt"
	"go-grpc-demo/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageSenderClient(conn)

	oneThreadTest(client)

	//fiveThreadsTest(client)
}

func oneThreadTest(client pb.MessageSenderClient) {
	startTime := time.Now()
	for i := 0; i < 100000; i++ {
		_, err := client.Send(context.Background(), &pb.MessageRequest{
			FirstNum:  32,
			SecondNum: 46,
		})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}
	endTime := time.Now()
	log.Println(fmt.Sprintf("oneThreadTest cost time: %.4f 秒", endTime.Sub(startTime).Seconds()))
}

func fiveThreadsTest(client pb.MessageSenderClient) {
	startTime := time.Now()
	wg := sync.WaitGroup{}
	for core := 0; core < 5; core++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 20000; i++ {
				_, err := client.Send(context.Background(), &pb.MessageRequest{
					FirstNum:  32,
					SecondNum: 46,
				})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
			}
		}()
	}
	wg.Wait()
	endTime := time.Now()
	log.Println(fmt.Sprintf("fiveThreadsTest cost time: %.4f 秒", endTime.Sub(startTime).Seconds()))
}
