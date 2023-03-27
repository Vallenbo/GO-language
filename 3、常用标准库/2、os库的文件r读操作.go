package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func main() {
	fileObj, err := os.Open(`F:\1\a.txt`) //打开一个文件，返回一个文件对象
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	//Bufioread(fileObj)
	//traditionRead(fileObj)
	OsRead()
	defer func(fileObj *os.File) { // 延迟关闭文件
		err := fileObj.Close()
		if err != nil {
			fmt.Println("close file failed!, err:", err)
		}
	}(fileObj)

}

func OsRead() { //3- ioutil.ReadFile读取整个文件
	content, err := os.ReadFile(`F:\1\a.txt`) //直接传入路径就能读取
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}
