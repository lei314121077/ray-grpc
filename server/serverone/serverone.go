package main

import (
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "pb/grpconepb"
	"sayapi"
)

const (
	port = ":50051"
)


func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGrpconepbServer(s, &sayapi.Server{})

	// 在gRPC服务器上注册反射服务。
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("未发现服务: %v", err)
	}

}