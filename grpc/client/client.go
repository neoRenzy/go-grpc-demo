package main

import (
	"context"
	"fmt"
	"github.com/shettyh/threadpool"
	"go-rpc-demo/pb"
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

	time.Sleep(time.Second * 10)

	threadPoolTest()
}

func NewGrpcClient() (client pb.GprcMessageSenderClient) {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client = pb.NewGRpcMessageSenderClient(conn)
	return
}

func oneThreadTest() {
	client := NewGrpcClient()
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
	log.Println(fmt.Sprintf("oneThreadTest cost time: %d 毫秒", endTime.Sub(startTime).Milliseconds()))
}

func fiveThreadsTest() {
	wg := sync.WaitGroup{}
	cntTime := make([]int64, 6)
	for core := 0; core < MAX_WORKER; core++ {
		wg.Add(1)
		go func(cnt int) {
			startTime := time.Now()
			client := NewGrpcClient()
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
			endTime := time.Now()
			cntTime[cnt] = endTime.Sub(startTime).Milliseconds()
		}(core)
	}
	wg.Wait()

	var all int64
	for _, t := range cntTime {
		all = all + t
	}
	log.Println(fmt.Sprintf("fiveThreadsTest cost time: %d 毫秒", all))
}

func threadPoolTest() {
	startTime := time.Now()

	threadPool := threadpool.NewThreadPool(MAX_WORKER, MAX_WORK_NUM)
	defer threadPool.Close()

	futures := make([]*threadpool.Future, MAX_WORK_NUM)
	for i := 0; i < MAX_WORK_NUM; i++ {
		future, _ := threadPool.ExecuteFuture(&XORNumTask{})
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

type XORNumTask struct{}

func (t XORNumTask) Call() interface{} {
	client := NewGrpcClient()
	defer client.Close()

	res, err := client.Send(context.Background(), &pb.MessageRequest{
		FirstNum:  FirstNum,
		SecondNum: SecondNum,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return res
}
