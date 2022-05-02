// @Author: Wilson 
// @Create: 2022/04/30 20:48:00
// @Description: TODO

package main

import (
	"GoMicroservices/gateway_demo-master/gateway_demo/demo/basic/unpack/unpack"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	unpack.Encode(conn, "hello world 0!!!")
}
