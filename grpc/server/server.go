package main

import (
	"context"
	"go-grpc-demo/grpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
)

type MessageSenderServerImpl struct {
	*pb.UnimplementedMessageSenderServer
}

func (MessageSenderServerImpl) Send(context context.Context, request *pb.MessageRequest) (*pb.MessageResponse, error) {
	log.Println("receive message:", request.GetSaySomething())
	resp := &pb.MessageResponse{}
	resp.ResponseSomething = "roger that!"
	return resp, nil
}

func main() {
	srv := grpc.NewServer()
	pb.RegisterMessageSenderServer(srv, MessageSenderServerImpl{})
	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
