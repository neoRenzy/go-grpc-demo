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
	"time"
)

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"

	"helloworld/api"
)

type GreeterProvider struct {
	api.UnimplementedMessageSenderServer
}

func (s *GreeterProvider) Send(ctx context.Context, request *api.MessageRequest) (*api.MessageResponse, error) {
	firstNum := request.GetFirstNum()
	secondNum := request.GetSecondNum()

	//log.Println("receive message: firstNum = ", firstNum)
	//log.Println("receive message: secondNum = ", secondNum)
	log.Println("receive message: time = ", time.Now())

	resp := &api.MessageResponse{
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
