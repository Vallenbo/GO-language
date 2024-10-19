package Concurrent

import (
	"context"
	"fmt"
	"time"
)

// var wg sync.WaitGroup

func worker(ctx context.Context) {
	go worker2(ctx) //调用worker2函数
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // io阻塞，等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) //返回具有“结束“命令的父级副本和cancel结束命令
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 3)
	cancel()  // 通知子goroutine结束命令
	wg.Wait() //等待所有登记的goroutine都结束
	fmt.Println("over")
}

func worker1() { // 1-1 始的例子
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	wg.Done() // 如何接收外部命令实现退出
}

func miantest() { //1-2 测试案例
	wg.Add(1)
	go worker1()
	wg.Wait() // 如何优雅的实现结束子goroutine
	fmt.Println("over")

}
