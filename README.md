# go-rpc-demo

## 从.proto文件生成.go文件
protoc --go_out=.  test.proto  
protoc --go-grpc_out=. test.proto

## 测试内容
基于整数XOR服务压测  
包含并发=1(串行压测)和并发=5两种，以及请求总数分别为10，100，1000，10000的测试
主要测试不同情况下单次请求的耗时

## 测试环境
### 硬件信息
芯片：	Apple M1 Pro  
核总数：	8（6性能和2能效）  
内存：	16 GB  

### 操作系统
macOS 12.3

### 语言
golang

### gRpc开源库信息
#### golang使用grpc方法
protoc --version  查看版本信息  
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip下载proto文件  
unzip protoc-3.15.8-linux-x86_64.zip -d /usr/.local/bin 解压到usr下的bin目录
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28 下载.proto文件生成.pb.go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2 下载.proto-->***_grpc.pb.go
#### 版本信息
google.golang.org/grpc v1.52.3
libprotoc v3.21.12

### dubbogo开源库信息
#### 版本信息
dubbo.apache.org/dubbo-go/v3 v3.0.1


## 压测结果
-|grpc-单线程|dubbo-单线程|grpc-五线程|dubbo-五线程|
|:---:|:---:|:---:|:---:|:---:|
|十次请求|||||
|百次请求|||||
|千次请求|||||
|万次请求|||||