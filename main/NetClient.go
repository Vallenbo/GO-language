package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000") //1-拨号连接到指定网络上的地址
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接

	reader := bufio.NewReader(os.Stdin)
	for {
		readString, _, _ := reader.ReadLine()
		fmt.Println("std read string ：", readString)
		conn.Write([]byte(readString))
		ConnReadString, _, _ := bufio.NewReader(conn).ReadLine()
		fmt.Println("Conn read string ：", string(ConnReadString))
	}
}

func NetClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:20000") //1-拨号连接到指定网络上的地址
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接

	buf := [512]byte{}
	bufioReader := bufio.NewReader(os.Stdin) //从终端发送数据
	for {
		input, _ := bufioReader.ReadString('\n') // 读取用户输入字符串
		inputInfo := strings.Trim(input, "\r\n") // 删除了 input 中所有的前导和尾随
		if strings.ToUpper(inputInfo) == "Q" {   // 如果输入q就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
