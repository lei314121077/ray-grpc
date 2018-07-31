# [grpc for go github](https://github.com/grpc/grpc-go)

* install

```bash
go get -u google.golang.org/grpc

# protobuf　package install　 
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
# protobuf 编译
protoc --go_out=plugins=grpc:. *.proto
``` 






