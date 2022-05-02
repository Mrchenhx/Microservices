// @Author: Wilson 
// @Create: 2022/04/30 20:47:00
// @Description: TODO

package main

import (
	"GoMicroservices/gateway_demo-master/gateway_demo/demo/basic/unpack/unpack"
	"fmt"
	"net"
)

func main() {
	// 1. 监听端口
	listener, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}
	// 2. 接收请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		bt, err := unpack.Decode(conn)
		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			break
		}
		str := string(bt)

		fmt.Printf("receive from client, data: %v\n", str)
	}
}
