#!/usr/bin/env bash
all:
	protoc -I/usr/local/include -I. \
		-I${HOME}/go/ray-grpc/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		--grpc-gateway_out=logtostderr=true:. \
		echoserver.proto
#all:
#	protoc -I/usr/local/include -I. \
#		-I${HOME}/go/ray-grpc/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#		--go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=$HOME/go/ray-grpc/src/github.com/golang/protobuf/protoc-gen-go/descriptor\
#		--grpc-gateway_out=logtostderr=true:. \
#		echoserver.proto

