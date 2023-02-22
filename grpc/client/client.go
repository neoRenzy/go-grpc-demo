package main

import (
	"context"
	"fmt"
	"github.com/shettyh/threadpool"
	"go-grpc-demo/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

const (
	MAX_WORK_NUM = 10000
	MAX_WORKER   = 5
	FirstNum     = 32
	SecondNum    = 46
)

func main() {
	oneThreadTest()
	time.Sleep(time.Second * 10)
	fiveThreadsTest()
	time.Sleep(time.Second * 10)
	threadPoolTest()
}

func NewClient() (client pb.MessageSenderClient) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client = pb.NewMessageSenderClient(conn)
	return
}

func oneThreadTest() {
	client := NewClient()

	startTime := time.Now()
	for i := 0; i < MAX_WORK_NUM; i++ {
		_, err := client.Send(context.Background(), &pb.MessageRequest{
			FirstNum:  FirstNum,
			SecondNum: SecondNum,
		})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
	}
	endTime := time.Now()
	log.Println(fmt.Sprintf("oneThreadTest cost time: %.4f 秒", endTime.Sub(startTime).Seconds()))
}

func fiveThreadsTest() {
	startTime := time.Now()
	wg := sync.WaitGroup{}
	for core := 0; core < MAX_WORKER; core++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			client := NewClient()
			for i := 0; i < MAX_WORK_NUM; i++ {
				_, err := client.Send(context.Background(), &pb.MessageRequest{
					FirstNum:  FirstNum,
					SecondNum: SecondNum,
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

func threadPoolTest() {
	startTime := time.Now()

	threadPool := threadpool.NewThreadPool(MAX_WORKER, MAX_WORK_NUM)
	defer threadPool.Close()

	futures := make([]*threadpool.Future, MAX_WORK_NUM)
	for i := 0; i < MAX_WORK_NUM; i++ {
		future, _ := threadPool.ExecuteFuture(&XORNumTask{
			client: NewClient(),
		})
		futures[i] = future
	}
	for {
		if futures[len(futures)-1].IsDone() {
			break
		}
	}

	endTime := time.Now()
	log.Println(fmt.Sprintf("threadPoolTest cost time: %.4f 秒", endTime.Sub(startTime).Seconds()))
}

type XORNumTask struct {
	client pb.MessageSenderClient
}

func (t XORNumTask) Call() interface{} {
	_, err := t.client.Send(context.Background(), &pb.MessageRequest{
		FirstNum:  FirstNum,
		SecondNum: SecondNum,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return nil
}
