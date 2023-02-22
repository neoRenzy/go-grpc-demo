package server

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	"go-grpc-demo/pb"
	"log"
	"time"
)

type GreeterProvider struct {
	pb.UnimplementedMessageSenderServer
}

func (GreeterProvider) Send(context context.Context, request *pb.MessageRequest) (*pb.MessageResponse, error) {
	firstNum := request.GetFirstNum()
	secondNum := request.GetSecondNum()

	//log.Println("receive message: firstNum = ", firstNum)
	//log.Println("receive message: secondNum = ", secondNum)
	log.Println("receive message: time = ", time.Now())

	resp := &pb.MessageResponse{
		Result: firstNum ^ secondNum,
	}
	return resp, nil
}

// export DUBBO_GO_CONFIG_PATH= PATH_TO_SAMPLES/helloworld/go-server/conf/dubbogo.yaml
func main() {
	config.SetProviderService(&GreeterProvider{})
	if err := config.Load(); err != nil {
		panic(err)
	}
	select {}
}
