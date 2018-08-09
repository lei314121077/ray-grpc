package main

import(
"net/rpc"
"log"
"fmt"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:64120") //64120为服务端启动服务的端口
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &Args{7,8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	// Asynchronous call
	clientTCP, err := rpc.Dial("tcp", "127.0.0.1:10011")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	quotient := new(Quotient)
	divCall := clientTCP.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done    // will be equal to divCall
	if replyCall.Error != nil {
		fmt.Println(replyCall.Error)
	} else {
		fmt.Printf("Arith: %d/%d=%d...%d\n", args.A, args.B, quotient.Quo, quotient.Rem)
	}
}
