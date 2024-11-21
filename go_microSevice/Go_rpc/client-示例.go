package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:8800") // 1. rpc 链接服务器 --Dial()
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer conn.Close()

	var reply string                                  // 接受返回值 --- 传出参数
	err = conn.Call("hello.HelloWorld", "李白", &reply) // 2. 调用远程函数
	if err != nil {
		fmt.Println("Call:", err)
		return
	}

	fmt.Println(reply) // 3、使用rpc的返回数据
}
