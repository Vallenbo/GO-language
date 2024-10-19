package Concurrent

import "fmt"

// channel声明
var ch1 chan<- int  // 声明一个传递整型的通道 //只能发送的通道
var ch2 <-chan bool // 声明一个传递布尔型的通道 //只能接收的通道
var ch3 chan []int  // 声明一个传递int切片的通道
func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func f2(ch chan int) {
	for v, ok := <-ch; ; { //循环取值
		if !ok {
			fmt.Println("通道已关闭")
			break
		}
		fmt.Printf("v:%#v ok:%#v\n", v, ok)
	}
}

func main() {
	ch := make(chan int, 2) // 1-声明一个缓冲区大小为2的通道
	ch <- 1
	ch <- 2
	close(ch)
	f2(ch)
	//unbufChannel() //无缓冲区通道
}

func unbufChannel() { //无缓冲区通道
	ch := make(chan int)
	go recv(ch) // 创建一个 goroutine 从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}

func f3(ch chan int) { //for range接收值
	for v := range ch {
		fmt.Println(v)
	}
}

func mutChannel() { //2- 多路复用
	ch := make(chan int, 1)
	for i := 1; i <= 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}
