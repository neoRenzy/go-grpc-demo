package main

import (
	"context"
	"go-grpc-demo/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

type MessageSenderServerImpl struct {
	*pb.UnimplementedMessageSenderServer
}

func (MessageSenderServerImpl) Send(context context.Context, request *pb.MessageRequest) (*pb.MessageResponse, error) {
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

func main() {
	// 单线程处理任务，port:8080
	oneThreadHandler()
	// 五线程处理任务，port:8081
	fiveThreadHandler()
}

func oneThreadHandler() {
	srv := grpc.NewServer()
	pb.RegisterMessageSenderServer(srv, MessageSenderServerImpl{})
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func fiveThreadHandler() {

}
