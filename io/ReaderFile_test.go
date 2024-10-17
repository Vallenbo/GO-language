package io

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func traditionRead(fileObj *os.File) { //1-传统文件读取方式
	var tmp [128]byte
	n, err := fileObj.Read(tmp[:])   // 将文件对象中的数据读取到，指定变量中
	if err != nil || err == io.EOF { //读取错误，或者读完了
		fmt.Println("read file failed!, err:", err)
		return
	}
	fmt.Printf("read word number is (byte):%d,and the content is :%v", n, string(tmp[:]))
}

func BufioRead(file *os.File) { //2- bufio是在file的基础上封装了一层API，支持更多的功能。
	reader := bufio.NewReader(file) //返回一个文件读取的对象
	for {
		line, err := reader.ReadString('\n') //注意是字符,指定读取截止符
		if err != nil || err == io.EOF {     //读取错误，或者读完了
			fmt.Println("read file failed!, err:", err)
			return
		}
		fmt.Print(line)
	}
}

func Test_readerFile(*testing.T) {
	//Bufioread(fileObj)
	//traditionRead(fileObj)
	content, err := os.ReadFile(`a.txt`) //直接传入路径就能读取
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
	//OsRead()

}

func openFile() *os.File {
	fileObj, err := os.Open(`F:\1\a.txt`) //打开一个文件，返回一个文件对象
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return fileObj
	}
	return fileObj
}
