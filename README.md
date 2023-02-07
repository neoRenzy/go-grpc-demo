# go-grpc-demo

##从.proto文件生成.go文件
protoc --go_out=.  test.proto  
protoc --go-grpc_out=. test.proto

##测试内容
基于整数XOR服务压测  
包含并发=1(串行压测)和并发=5两种

##测试环境
###硬件信息
芯片：	Apple M1 Pro  
核总数：	8（6性能和2能效）  
内存：	16 GB  

###操作系统
macOS 12.3

###语言
golang

###gRpc开源库信息
####golang使用grpc方法
protoc --version  查看版本信息  
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip下载proto文件  
unzip protoc-3.15.8-linux-x86_64.zip -d /usr/.local/bin 解压到usr下的bin目录
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 下载.proto文件生成.pb.go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 下载.proto-->***_grpc.pb.go
####版本信息
google.golang.org/grpc v1.52.3
libprotoc v3.21.12

###thrift信息

## 压测结果
同为十万次请求，分为单协程处理和五协程并发处理两种情况
###grpc测试结果
串行处理时间 : 6.0399 秒
并发处理时间: 9.7539 秒