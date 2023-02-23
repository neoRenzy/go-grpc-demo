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
	"log"
)

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"go-rpc-demo/dubbo/api"
)

// export DUBBO_GO_CONFIG_PATH=./go-client/conf/dubbogo.yaml
func main() {
	grpcGreeterImpl := &api.MessageSenderClientImpl{}
	grpcGreeterImpl2 := &api.MessageSenderClientImpl2{}
	grpcGreeterImpl3 := &api.MessageSenderClientImpl3{}
	grpcGreeterImpl4 := &api.MessageSenderClientImpl4{}
	grpcGreeterImpl5 := &api.MessageSenderClientImpl5{}
	config.SetConsumerService(grpcGreeterImpl)
	config.SetConsumerService(grpcGreeterImpl2)
	config.SetConsumerService(grpcGreeterImpl3)
	config.SetConsumerService(grpcGreeterImpl4)
	config.SetConsumerService(grpcGreeterImpl5)
	if err := config.Load(); err != nil {
		panic(err)
	}

	log.Print("start to test dubbo")
	req := &api.MessageRequest{
		FirstNum:  11,
		SecondNum: 12,
	}
	ctx := context.Background()
	for i := 0; i < 5; i++ {
		reply, err := grpcGreeterImpl.Send(ctx, req)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("client response result: %v\n,time : %d", reply, i)
	}
	//reply, err = grpcGreeterImpl2.Send(context.Background(), req)
	//if err != nil {
	//	logger.Error(err)
	//}
	//logger.Infof("client2 response result: %v\n", reply)
}
