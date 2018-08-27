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

## “there are no silver bullets”

* [微服务架构文章参考](http://microservices.io/)
        
* [微服务架构文章参考](https://studygolang.com/search?q=%E5%BE%AE%E6%9C%8D%E5%8A%A1%E5%AE%9E%E6%88%98)
    
    * LEVE 1 [微服务架构的优势与不足](https://studygolang.com/articles/6214)
    
    ```
    单体式的架构更适合轻量级的简单应用。如果你用它来开发复杂应用，那真的会很糟糕。
    微服务架构模式可以用来构建复杂应用，当然，这种架构模型也有自己的缺点和挑战。
    ```
    
    * LEVE 2 [apigateway](https://studygolang.com/articles/6230) 
      
       对于大多数微服务基础的应用，实现一个API Gateway都是有意义的，它就像是进入系统的一个服务提供点。
       API Gateway负责请求转发、请求合成和协议转换。它提供给应用客户端一个自定义的API。
       API Gateway可以通过返回缓存或者默认值的方式来掩盖后端服务的错误。
    
    *  LEVE 3 [深入微服务架构的进程间通信](https://studygolang.com/articles/6249) 
    
        * 交互模式
        
        当为某一个服务选择IPC时，首先需要考虑服务之间如何交互。客户端和服务器之间有很多的交互模式，我们可以从两个维度进行归类。第一个维度是一对一还是一对多：
        
            • 一对一：每个客户端请求有一个服务实例来响应。
            • 一对多：每个客户端请求有多个服务实例来响应
            
        第二个维度是这些交互式同步还是异步：
        
            • 同步模式：客户端请求需要服务端即时响应，甚至可能由于等待而阻塞。
            • 异步模式：客户端请求不会阻塞进程，服务端的响应可以是非即时的。
            
        * 一对一的交互模式有以下几种方式：
        
            • 请求/响应：
            
            ```
            一个客户端向服务器端发起请求，等待响应。客户端期望此响应即时到达。在一个基于线程的应用中，等待过程可能造成线程阻塞。
            ```
            
            • 通知（也就是常说的单向请求）：
            
            ```
            一个客户端请求发送到服务端，但是并不期望服务端响应。
            ```
            
            • 请求/异步响应：
            
            ```
            客户端发送请求到服务端，服务端异步响应请求。客户端不会阻塞，而且被设计成默认响应不会立刻到达。
            ```  
        * 一对多的交互模式有以下几种方式：
        
            • 发布/ 订阅模式：
            
            ```
            客户端发布通知消息，被零个或者多个感兴趣的服务消费。
            ```
            
            • 发布/异步响应模式：
            
            ```
            客户端发布请求消息，然后等待从感兴趣服务发回的响应。
            ```
    
## 关于服务器容错机制熔断和降级


* 服务降级：

    ```
      主逻辑失败采用备用逻辑的过程
      服务降级，就是整体资源快不够了，忍痛将某些服务先关掉，待渡过难关，再开启回来。
      在调用A服务失败的时候达到预期能够接受的最大值的时候，调用B服务，满足要求，每次都是这个顺序 
    ```
    
* 服务熔断：又叫做过载保护

    ```
    主逻辑多次失败采用备用逻辑的过程，暂时性忽略主逻辑并不再采用，而采用备用逻辑的过程
    服务降级，就是整体资源快不够用了，将某些服务先关掉，一段时间内切换到备用的服务方案，再开启回来。
    在调用A服务失败的时候达到预期能够接受的最大值的时候，调用B服务，以后一定时间段内就一直调用B服务了
    ```
    
* [服务熔断机制hystrix](https://blog.csdn.net/weixin_40584932/article/details/80903772)



### 一些分布式微服务开发中的经验和策略

*  关于部分失败造成的用户体验差问题

   • 网络超时：
   
   ```
   当等待响应时，不要无限期的阻塞，而是采用超时策略。使用超时策略可以确保资源不会无限期的占用。
   ```
   
   • 限制请求的次数：
   
   ```
   可以为客户端对某特定服务的请求设置一个访问上限。如果请求已达上限，就要立刻终止请求服务。
   ```
   
   • 断路器模式（Circuit Breaker Pattern）：
   
   ```   
   记录成功和失败请求的数量。如果失效率超过一个阈值，触发断路器使得后续的请求立刻失败。如果大量的请求失败，就可能是这个服务不可用，再发请求也无意义。在一个失效期后，客户端可以再试，如果成功，关闭此断路器。
   ```
   
   • 提供回滚：
   
   ```
   当一个请求失败后可以进行回滚逻辑。例如，返回缓存数据或者一个系统默认值。
   ```

