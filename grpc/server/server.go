package main

import (
	"context"
	"fmt"
	"go-grpc-demo/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func main() {
	// port:8080
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

	fmt.Println("server start")
}

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
