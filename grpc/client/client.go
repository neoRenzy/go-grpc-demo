package main

import (
	"context"
	"fmt"
	"github.com/shettyh/threadpool"
	"go-grpc-demo/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	//time.Sleep(time.Second * 10)
	//threadPoolTest()
}

func NewClient() (client pb.GprcMessageSenderClient) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = pb.NewGRpcMessageSenderClient(conn)
	return
}

func oneThreadTest() {
	client := NewClient()
	defer client.Close()

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
			client := NewClient()
			defer wg.Done()
			defer client.Close()

			for i := 0; i < MAX_WORK_NUM/MAX_WORKER; i++ {
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
		client := NewClient()

		future, _ := threadPool.ExecuteFuture(&XORNumTask{
			client: client,
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
	client pb.GprcMessageSenderClient
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
