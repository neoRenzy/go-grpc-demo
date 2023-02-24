package main

import (
	"context"
	"fmt"
	"github.com/shettyh/threadpool"
	"go-rpc-demo/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"math"
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

	fmt.Println("one thread test complete")
	time.Sleep(time.Second * 3)

	fiveThreadsTest()

	//threadPoolTest()
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
	log.Println(fmt.Sprintf("oneThreadTest cost time: %d 微秒", endTime.Sub(startTime).Microseconds()/MAX_WORK_NUM))
}

func fiveThreadsTest() {
	wg := sync.WaitGroup{}
	cntTime := make([]int64, 5)
	for core := 0; core < MAX_WORKER; core++ {
		wg.Add(1)
		go func(cnt int) {
			client := NewGrpcClient()
			defer wg.Done()
			defer client.Close()
			// 去除新建连接的时间
			startTime := time.Now()
			for i := 0; i < MAX_WORK_NUM/MAX_WORKER; i++ {
				_, err := client.Send(context.Background(), &pb.MessageRequest{
					FirstNum:  FirstNum,
					SecondNum: SecondNum,
				})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
			}
			cntTime[cnt] = time.Now().Sub(startTime).Microseconds()
		}(core)
	}
	wg.Wait()

	var all float64
	for _, t := range cntTime {
		all = math.Max(float64(t), all)
	}
	log.Println(fmt.Sprintf("fiveThreadsTest cost time: %.1f 微秒", all/MAX_WORK_NUM))
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
