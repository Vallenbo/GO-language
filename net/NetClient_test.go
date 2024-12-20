package net

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
)

func Test_NetClient(t *testing.T) { // client端
	conn, err := net.Dial("tcp", "127.0.0.1:20000") //1-拨号连接到指定网络上的地址
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close() // 关闭连接
	conn.RemoteAddr().String()
	bufioReader := bufio.NewReader(os.Stdin) //从终端发送数据
	for {
		input, _ := bufioReader.ReadString('\n') // 读取用户输入字符串
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func Test_NetClientBufioReader(t *testing.T) { // client端
	conn, err := net.Dial("tcp", "127.0.0.1:20000") //1-拨号连接到指定网络上的地址
	if err != nil {
		fmt.Println("err :", err)
		return
	}
	defer conn.Close()                       // 关闭连接
	bufioReader := bufio.NewReader(os.Stdin) //从终端发送数据
	bufioWriter := bufio.NewWriter(conn)
	for {
		input, _ := bufioReader.ReadString('\n') // 读取用户输入字符串
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" { // 如果输入q就退出
			return
		}
		_, err = bufioWriter.WriteString(inputInfo) // 发送数据
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed, err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
