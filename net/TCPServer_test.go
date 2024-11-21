package net

import (
	"log"
	"net"
	"testing"
)

func Test_TCPClient(t *testing.T) {
	// 解析服务端监听地址，本例以tcp为例
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Panic(err)
	}
	// 创建监听器
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Panic(err)
	}
	for {
		// 监听客户端连接请求
		conn, err := listen.AcceptTCP()
		if err != nil {
			continue
		}
		// 处理客户端请求 这个函数可以自己编写

		go process(conn)
	}
}
