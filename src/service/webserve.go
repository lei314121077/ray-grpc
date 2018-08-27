package service

import (
	"fmt"
	"log"
	"net/http"
)

//@name StartWebServer
//@param port string 端口
func StartWebServer(addr, port string) {
	r := NewRouter()             // NEW
	http.Handle("/", r)          // NEW

	log.Println("启动http服务 " + port)
	err := http.ListenAndServe(fmt.Sprintf("%v:%v", addr, port) , nil)    // Goroutine will block here

	if err != nil {
		log.Println("端口启动HTTP侦听器时发生错误 " + port)
		log.Println("Error: " + err.Error())
	}
}
