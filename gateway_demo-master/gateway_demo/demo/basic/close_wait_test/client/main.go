/**
  @Author: Wilson
  @Create: 2022/04/30 17:14:00
  @Description: TODO
 **/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 模拟发起请求
	doSend()
	fmt.Println("doSend over")
	doSend()
	fmt.Println("doSend over")
}

func doSend() {
	// 1. 连接服务器
	conn, err := net.Dial("tcp", "localhost:9090")
	//defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err: %v\n", err)
		return
	}
	// 2. 读取命令行输入
	inputReader := bufio.NewReader(os.Stdin) // bufio 内有一个 buffer，可以缓冲数据
	for {
		// 3. 读取数据直到遇到 \n
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("read from console failed, err: %v\n", err)
			break
		}
		// 4. 读取到 Q 时停止
		trimmedInput := strings.TrimSpace(input)
		if trimmedInput == "Q" {
			break
		}
		// 5. 回复服务器消息
		_, err = conn.Write([]byte(trimmedInput))
		if err != nil {
			fmt.Printf("write failed, err: %v\n", err)
			break
		}
	}

}