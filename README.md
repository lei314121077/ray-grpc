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
..
 ├── client
 │   └── clientone
 │       ├── build.sh
 │       ├── clientone.go
 │       ├── docker-compose.yml
 │       ├── Dockerfile
 │       └── src
 │           └── saymsg
 │               └── saymsg.go
 ├── main.go
 ├── README.md
 ├── server
 │   ├── serverone
 │   │   ├── build.sh
 │   │   ├── docker-compose.yml
 │   │   ├── Dockerfile
 │   │   ├── serverone.go
 │   │   └── src
 │   │       └── sayapi
 │   │           ├── sayapi.go
 │   │           ├── say.go
 │   │           └── say_test.go
 │   └── servertwo
 │       ├── build.sh
 │       ├── docker-compose.yml
 │       ├── Dockerfile
 │       ├── servertwo.go
 │       └── src
 │           └── echoapi
 │               └── echoapi.go
 └── src
     ├── pb
     │   ├── echoserver
     │   │   ├── build.sh
     │   │   ├── echoserver.pb.go
     │   │   ├── echoserver.pb.gw.go
     │   │   ├── echoserver.proto
     │   │   ├── gen.go
     │   │   └── Makefile
     │   └── grpconepb
     │       ├── build.sh
     │       ├── grpconepb.pb.go
     │       └── grpconepb.proto
     └── ray
         ├── common
         │   └── common.go
         └── gateway
             └── wsproxy.go
 



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

## 资料

* [微服务架构文章参考](http://microservices.io/)
* [微服务架构文章参考](https://studygolang.com/search?q=%E5%BE%AE%E6%9C%8D%E5%8A%A1%E5%AE%9E%E6%88%98)

## [apigateway](https://studygolang.com/articles/6230) 的形似

 对于大多数微服务基础的应用，实现一个API Gateway都是有意义的，它就像是进入系统的一个服务提供点。
 API Gateway负责请求转发、请求合成和协议转换。它提供给应用客户端一个自定义的API。
 API Gateway可以通过返回缓存或者默认值的方式来掩盖后端服务的错误。


