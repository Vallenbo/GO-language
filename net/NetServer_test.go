package net

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"
)

func Test_NetServer(t *testing.T) { //server端
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
		fmt.Println("a new Accept success.")
		go process(conn) // 启动一个goroutine处理连接
	}
}

func process(conn net.Conn) { // 处理函数
	defer conn.Close() // 关闭连接
	var buf [128]byte
	//n, err := conn.Read(buf[:])   //从连接中读取数据
	reader := bufio.NewReader(conn)
	for {
		n, err := reader.Read(buf[:]) // 读取数据
		//n, isPrefix, err := reader.ReadLine() // 读取数据
		//if isPrefix || err != nil {
		//	fmt.Println("read line is end, isPrefix:", isPrefix)
		//	fmt.Println("read from client failed, err:", err)
		//	break
		//}
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", recvStr)
		time.Sleep(time.Second)
		conn.Write([]byte(recvStr)) // 发送数据
	}
}
