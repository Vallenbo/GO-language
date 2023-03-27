package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func hello() {
	fmt.Println("hello")
}
func hello1(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("hello", i)
}

func individualGo() { //单个goroutine
	go hello()              //开启一个单独的goroutine用户协程，去执行hello()函数
	fmt.Println("你好")     //由于main函数启动的goroutine结束了
	time.Sleep(time.Second) //进行等待
}
func multipleGo() {
	multipleGo()
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello1(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}
func main() {
	//runtime.GOMAXPROCS(2) //设置cpu数

}
