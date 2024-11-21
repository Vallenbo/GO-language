package main

import (
	"Go_grpc/hello_grpc"
	"fmt"
	"io"
	"net/http"
)

// gin的grpc客户端 ， 用来测试gin框架返回的protobuf是否正确
func main() {
	addr := ":8080"
	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
	// 此处使用不安全的证书来实现 SSL/TLS 连接
	resp, err := http.Get("http://localhost:8080/GetProtoBuf")
	if err != nil {
		panic(fmt.Sprintf("http connect addr [%s] 连接失败 %s", addr, err))
	}
	defer resp.Body.Close()
	// 初始化客户端

	body, err := io.ReadAll(resp.Body) //读取所有内容
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}

	test := hello_grpc.Test{
		Msg: string(body),
	}

	fmt.Println("test.Msg:", test.Msg)
}
