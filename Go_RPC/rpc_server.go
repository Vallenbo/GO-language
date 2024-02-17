package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type World struct {
}

func (*World) HelloWorld(name string, resp *string) error {
	*resp = name + " 你好！"
	//return nil
	return errors.New("未知的错误！")
}

func main() {
	err := rpc.RegisterName("Hello", new(World))
	if err != nil {
		fmt.Println("rpc RegisterName is error!")
		return
	}

	listen, err := net.Listen("tcp", "127.0.0.1:8800")
	defer listen.Close()

	fmt.Println("开始监听。。。。")
	accept, err := listen.Accept()
	if err != nil {
		fmt.Println("rpc Accept is error!")
		return
	}
	defer accept.Close()

	rpc.ServeConn(accept)
	//jsonrpc.ServeConn(conn) // jsonrpc数据序列化

}
