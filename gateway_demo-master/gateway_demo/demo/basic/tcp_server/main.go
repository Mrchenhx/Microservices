/**
  @Author: Wilson
  @Create: 2022/04/30 17:14:00
  @Description: close_wait_test 测试
 **/
package main

import (
	"fmt"
	"net"
)

func main() {
	//1、监听端口
	listen, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		fmt.Printf("listen fail, err: %v\n", err)
		return
	}
	fmt.Println("Server Run in 0.0.0.0:9090")
	//2.建立套接字连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept fail, err: %v\n", err)
			continue // 监听下一个连接
		}
		//3. 创建处理协程
		go process(conn)
	}
}

func process(conn net.Conn){
	//defer conn.Close() // 不关闭会出现 close_wait 状态
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read from connect failed, err: %v\n", err)
			// 客户端主动关闭，若没有 conn.Close() 则服务器会出现 close_wait，
			// 客户端则出现 Fin_wait
			break
		}
		str := string(buf[:n])
		fmt.Printf("receive from client, data: %v\n", str)
	}
}