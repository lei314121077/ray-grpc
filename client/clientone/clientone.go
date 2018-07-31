package main

import (
	"log"
	"os"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "pb/grpconepb"
)

const (
	address     = "localhost:50051"
	defaultName = "kitty"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("未发现链接!: %v", err)
	}
	defer conn.Close()
	c := pb.NewGrpconepbClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHi(ctx, &pb.Say{Name: name, Msg:"hello world!"})
	if err != nil {
		log.Fatalf("未响应!: %v", err)
	}
	log.Printf("Greeting: %s", r.Body)
}



