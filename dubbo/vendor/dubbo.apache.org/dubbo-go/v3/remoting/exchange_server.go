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

package remoting

import (
	"dubbo.apache.org/dubbo-go/v3/common"
)

// Server is the interface that wraps the basic Start method and Stop method.
// It is interface of server for network communication. If you use getty as network
// communication, you should define GettyServer that implements this interface.
//
// Start method invokes once for connection.
//
// Stop method is for destroy.
type Server interface {
	Start()
	Stop()
}

// ExchangeServer is abstraction level. it is like facade. it implements Start and Stop.
type ExchangeServer struct {
	Server Server
	URL    *common.URL
}

// NewExchangeServer returns a ExchangeServer that constructs from url and server.
func NewExchangeServer(url *common.URL, server Server) *ExchangeServer {
	exchangServer := &ExchangeServer{
		Server: server,
		URL:    url,
	}
	return exchangServer
}

func (server *ExchangeServer) Start() {
	server.Server.Start()
}

func (server *ExchangeServer) Stop() {
	server.Server.Stop()
}
