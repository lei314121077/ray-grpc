# [grpc for go github](https://github.com/grpc/grpc-go)

[![grpc-gateway](https://travis-ci.org/grpc-ecosystem/grpc-gateway.svg?branch=master)](#https://github.com/lei314121077/grpc-gateway)
[![servertwo Build Status](https://travis-ci.org/grpc-ecosystem/grpc-gateway.svg?branch=master)](https://travis-ci.org/grpc-ecosystem/grpc-gateway)
[![参考](https://godoc.org/github.com/tmc/grpc-websocket-proxy/wsproxy)](https://github.com/tmc/grpc-websocket-proxy)

* servertwo 支持httph和tcp websocket 协议

   ![architecture introduction diagram](https://docs.google.com/drawings/d/12hp4CPqrNPFhattL_cIoJptFvlAqm5wLQ0ggqI5mkCg/pub?w=749&amp;h=370)

* [GRPC概述](https://grpc.io/docs/quickstart/go.html)

gRPC 是一个高性能、通用的开源RPC框架，其由 Google 主要面向移动应用开发并基于HTTP/2 协议标准而设计，
基于 ProtoBuf(Protocol Buffers) 序列化协议开发，且支持众多开发语言。
gRPC 的基石就是 HTTP/2，然后在上面使用 protobuf 协议定义好 service RPC.


* 项目结构

```bash
.
├── client
│   └── clientone
│       ├── build.sh
│       ├── clientone.go
│       ├── docker-compose.yml
│       └── Dockerfile
├── main.go
├── README.md
├── server
│   └── serverone
│       ├── build.sh
│       ├── docker-compose.yml
│       ├── Dockerfile
│       ├── serverone.go
│       └── src
│           └── sayapi
│               ├── sayapi.go
│               ├── say.go
│               └── say_test.go
└── src            # 公共目录,为了方便就都写在一个里面了
    ├── pb         
    │   └── grpconepb
    │       ├── build.sh
    │       ├── grpconepb.pb.go
    │       └── grpconepb.proto
    └── ray        # 个人工具包
        └── common
            └── common.go


```


* install

```bash
go get -u google.golang.org/grpc

# protobuf　package install　 
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
# protobuf 编译
protoc --go_out=plugins=grpc:. *.proto
``` 


* 关于protobuf3的接口类型定义

```
gRPC 定义了四种类型的服务接口:
一元RPC，客户端向服务器发送请求并获得响应，就像正常的函数调用一样。
rpc SayHello(HelloRequest) returns (HelloResponse){
}
服务器流RPC，客户端发送一个对象服务器端返回一个Stream（流式消息）
rpc LotsOfReplies(HelloRequest) returns (stream HelloResponse){
}
客户端流式RPC，客户端发送一个Stream（流式消息）服务端返回一个对象。
rpc LotsOfGreetings(stream HelloRequest) returns (HelloResponse) {
}
```




