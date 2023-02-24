/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"go-rpc-demo/dubbo/api"
)

const (
	MAX_WORK_NUM = 100000
	MAX_WORKER   = 5
	FirstNum     = 32
	SecondNum    = 46
)

var grpcGreeterImpl = &api.MessageSenderClientImpl{}
var grpcGreeterImpl2 = &api.MessageSenderClientImpl2{}
var grpcGreeterImpl3 = &api.MessageSenderClientImpl3{}
var grpcGreeterImpl4 = &api.MessageSenderClientImpl4{}
var grpcGreeterImpl5 = &api.MessageSenderClientImpl5{}

// export DUBBO_GO_CONFIG_PATH=./go-client/conf/dubbogo.yaml
func main() {
	config.SetConsumerService(grpcGreeterImpl)
	config.SetConsumerService(grpcGreeterImpl2)
	config.SetConsumerService(grpcGreeterImpl3)
	config.SetConsumerService(grpcGreeterImpl4)
	config.SetConsumerService(grpcGreeterImpl5)
	if err := config.Load(); err != nil {
		panic(err)
	}
	log.Print("start to test dubbo")

	oneThreadTest()

	time.Sleep(2 * time.Second)

	fiveThreadTest()
}

func oneThreadTest() {
	req := &api.MessageRequest{
		FirstNum:  FirstNum,
		SecondNum: SecondNum,
	}
	startTime := time.Now()
	for i := 0; i < MAX_WORK_NUM; i++ {
		_, _ = grpcGreeterImpl.Send(context.Background(), req)
	}
	fmt.Printf("oneThreadTest cost time: %d 微秒", time.Now().Sub(startTime).Microseconds()/MAX_WORK_NUM)
}

func fiveThreadTest() {
	req := &api.MessageRequest{
		FirstNum:  FirstNum,
		SecondNum: SecondNum,
	}
	ctx := context.Background()

	wg := sync.WaitGroup{}
	wg.Add(5)

	startTime := time.Now()
	go func() {
		defer wg.Done()
		for i := 0; i < MAX_WORK_NUM/MAX_WORKER; i++ {
			_, _ = grpcGreeterImpl.Send(ctx, req)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < MAX_WORK_NUM/MAX_WORKER; i++ {
			_, _ = grpcGreeterImpl2.Send(ctx, req)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < MAX_WORK_NUM/MAX_WORKER; i++ {
			_, _ = grpcGreeterImpl3.Send(ctx, req)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < MAX_WORK_NUM/MAX_WORKER; i++ {
			_, _ = grpcGreeterImpl4.Send(ctx, req)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < MAX_WORK_NUM/MAX_WORKER; i++ {
			_, _ = grpcGreeterImpl5.Send(ctx, req)
		}
	}()

	wg.Wait()
	fmt.Printf("fiveThread cost time: %d 微秒", time.Now().Sub(startTime).Microseconds()/MAX_WORK_NUM)
}
