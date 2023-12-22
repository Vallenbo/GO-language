package main

import (
	"bufio"
	"fmt"
	"net"
)

// TCP server端
func process(conn net.Conn) { // 处理函数
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn) // 关闭连接
	for {
		var buf [128]byte
		//n, err := conn.Read(buf[:])   //从连接中读取数据
		reader := bufio.NewReader(conn)

		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}

func main() { //server端
	listenObj, err := net.Listen("tcp", "127.0.0.1:20000") //1-监听本地网络地址上的通知，返回监听对象
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for { //循环等待
		conn, err := listenObj.Accept() // 2-等待接受，并返回一个连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}
