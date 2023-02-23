package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	"go-grpc-demo/pb"
	"log"
)

const (
	MAX_WORK_NUM = 10000
	MAX_WORKER   = 5
	FirstNum     = 32
	SecondNum    = 46
)

// export DUBBO_GO_CONFIG_PATH= PATH_TO_SAMPLES/helloworld/go-client/conf/dubbogo.yaml
func main() {
	var grpcGreeterImpl = new(pb.MessageSenderClientImpl)

	config.SetConsumerService(grpcGreeterImpl)
	if err := config.Load(); err != nil {
		panic(err)
	}

	log.Printf("start to test dubbo")
	reply, err := grpcGreeterImpl.Send(context.Background(), &pb.MessageRequest{
		FirstNum:  FirstNum,
		SecondNum: SecondNum,
	})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	log.Printf("client response result: %v\n", reply)
}
