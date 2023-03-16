package main

import (
	"bufio"
	"fmt"
	"os"
)

func traditionWrite(fileObj *os.File) { //1-传统的写入方式 Write和WriteString
	str := "hello 沙河"
	fileObj.Seek(1, 0)               //设置文件读写位置，offset偏移量,whencew位置
	fileObj.Write([]byte(str))       //写入字节切片数据
	fileObj.WriteString("hello 小王子") //直接写入字符串数据
}

func bufioWrite(fileObj *os.File) { //2-bufio.NewWriter ，可循环写入
	writer := bufio.NewWriter(fileObj) //返回一个可写的操作对象
	//writer = bufio.NewWriter(os.Stdin) //从键盘输入
	for i := 0; i < 10; i++ {
		writer.WriteString("hello沙河\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}

func main() {
	fileObj, err := os.OpenFile(`F:\1\a.txt`, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644) //或者追加，或者创建
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	traditionWrite(fileObj)
	//bufioWrite(fileObj)
	//OslWrite(fileObj)
	defer func(fileObj *os.File) { // 延迟关闭文件
		err := fileObj.Close()
		if err != nil {
			fmt.Println("close file failed!, err:", err)
		}
	}(fileObj)
}

func OslWrite(fileObj *os.File) { //os.WriteFile 可直接写入
	str := "hello 沙河"
	err := os.WriteFile("./xx.txt", []byte(str), 0666) //写入文件将数据写入命名文件，并在必要时创建它。
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
