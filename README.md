# [grpc for go github](https://github.com/grpc/grpc-go)

[![API](#)](#)
[![BEANCHMARK](#)](#)
[![PACKAGE](#)](#)
[![DOC](#)](#)

* 概述[参见](https://www.jianshu.com/p/d27ac8039fff)

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






