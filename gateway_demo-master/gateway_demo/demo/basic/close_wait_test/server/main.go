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
		go func(conn net.Conn) {
			//defer conn.Close() // 不关闭会出现 close_wait 状态
			for {
				var buf [128]byte
				n, err := conn.Read(buf[:])
				if err != nil {
					fmt.Printf("read from connect failed, err: %v\n", err)
					break // 读取数据失败，应该停止当前连接，执行 defer 语句
				}
				str := string(buf[:n])
				fmt.Printf("receive from client, data: %v\n", str)
			}
		}(conn)
	}
}
