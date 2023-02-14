package main

import (
	"context"
	"fmt"
	"github.com/shettyh/threadpool"
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
	// 五线程处理任务，port:8081
	fiveThreadHandler()
	// 单线程处理任务，port:8080
	oneThreadHandler()

	fmt.Println("server start")
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

type MessageSenderServerConcurrencyImpl struct {
	Pool *threadpool.ThreadPool
	*pb.UnimplementedMessageSenderServer
}

type taskSend struct {
	in *pb.MessageRequest
}

func (t *taskSend) Call() interface{} {
	firstNum := t.in.FirstNum
	secondNum := t.in.SecondNum

	return firstNum ^ secondNum
}

func (p MessageSenderServerConcurrencyImpl) Send(context context.Context, request *pb.MessageRequest) (*pb.MessageResponse, error) {
	task := &taskSend{
		in: request,
	}
	future, err := p.Pool.ExecuteFuture(task)
	if err != nil {
		return nil, err
	}

	for {
		if future.IsDone() {
			break
		}
	}

	resp := &pb.MessageResponse{
		Result: future.Get().(int32),
	}
	return resp, nil
}

func fiveThreadHandler() {
	pool := threadpool.NewThreadPool(5, 20000)

	srv := grpc.NewServer()
	pb.RegisterConcurrencyMessageSenderServer(srv, MessageSenderServerConcurrencyImpl{
		Pool: pool,
	})
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
