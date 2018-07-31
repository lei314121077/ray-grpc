# [grpc for go github](https://github.com/grpc/grpc-go)

[![API](#)](#)
[![BEANCHMARK](#)](#)
[![PACKAGE](#)](#)
[![DOC](#)](#)


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






