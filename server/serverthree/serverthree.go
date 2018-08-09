package main

import (
	"flag"
	"log"
	"net/rpc"
	"net"
	"sync"
	"api"
	"net/http/httptest"
	"net/http"
)

var (
	httpaddr = flag.String("httpaddr", "localhost:8972", "http server address")
	tcpaddr = flag.String("tcpaddr", "localhost:10011", "tcp server addrdss")
)


func listenTCP() (net.Listener, string) {
	l, e := net.Listen("tcp", *tcpaddr)
	if e != nil {
		log.Fatalf("net.Listen tcp :0: %v", e)
	}
	return l, l.Addr().String()
}

func startHttpServer() {
	//case 1
	//http.HandleFunc("/hello", api.SayHello)
	//server := http.ListenAndServe(*httpaddr, nil)

	//case 2
	server := httptest.NewServer(nil)
	//server := httptest.NewTLSServer(nil)
	httpServerAddr := server.Listener.Addr().String()
	log.Println("Test HTTP RPC server listening on", httpServerAddr)

	//case 3
	////body := "{\"action\":20}"
	//res, err := http.Post(*httpaddr, "application/json;charset=utf-8", bytes.NewBuffer([]byte(pb.body)))
	//if err != nil {
	//	fmt.Println("Fatal error ", err.Error())
	//}
	//
	//defer res.Body.Close()
	//
	//content, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println("Fatal error ", err.Error())
	//}
	//
	//fmt.Println(string(content))
}

func main() {

	rpc.Register(new(api.Arith)) //注册服务
	var l net.Listener
	l, serverAddr := listenTCP() //监听TCP连接
	log.Println("Test RPC server listening on", serverAddr)
	go rpc.Accept(l)

	rpc.HandleHTTP() //监听HTTP连接

	var httpOnce sync.Once
	httpOnce.Do(startHttpServer)

	select{}
}




