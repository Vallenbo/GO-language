package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) { // cat命令实现
	for {
		buf, err := r.ReadBytes('\n') //注意是字符
		if err == io.EOF {
			fmt.Fprintf(os.Stdout, "%s", buf) // 退出之前将已读到的内容输出
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}

func main() {
	flag.Parse() // 解析命令行参数
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin)) // 如果没有参数默认从标准输入读取内容
	}

	for i := 0; i < flag.NArg(); i++ { // 依次读取每个指定文件的内容并打印到终端
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err:%v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}
