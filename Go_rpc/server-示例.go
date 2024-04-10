package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type World struct { // 定义类对象
}

func (this *World) HelloWorld(name string, resp *string) error { // 绑定类方法
	*resp = name + " 你好!"
	return nil
}

func main() {
	err := rpc.RegisterName("hello", new(World)) // 1. rpc, 注册RPC服务, 绑定对象方法（rpc格式的类方法）
	if err != nil {
		fmt.Println("注册 rpc 服务失败!", err)
		return
	}

	listener, err := net.Listen("tcp", "127.0.0.1:8800") // 2. net,设置tcp监听器
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer listener.Close()

	fmt.Println("开始监听 ...")
	conn, err := listener.Accept() // 3. net,建立链接
	if err != nil {
		fmt.Println("Accept() err:", err)
		return
	}
	defer conn.Close()
	fmt.Println("链接成功...")

	rpc.ServeConn(conn) // 4. net，rpc，绑定服务
}
