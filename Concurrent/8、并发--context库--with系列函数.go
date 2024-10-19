package Concurrent

import (
	"context"
	"fmt"
	"time"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func WithCancelFunc() { //1- WithCancel函数使用
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们取完需要的整数后调用cancel

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
func WithDeadlineFunc() { //2- WithDeadline函数使用
	d := time.Now().Add(5 * time.Second)                         //指定截止时间
	ctx, cancel := context.WithDeadline(context.Background(), d) //返回具有“结束“命令和截止时间的父级副本和cancel结束命令

	defer cancel() //官方建议，尽量最后，延迟调用cancel函数，减少不必要的运行时间

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

//var wg sync.WaitGroup

func worker3(ctx context.Context) { // 3-1 示例函数
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func WithTimeoutFunc() { // 3-2 WithTimeout函数使用
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50) // 设置一个50毫秒的超时
	wg.Add(1)
	go worker3(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}

type TraceCode string

func worker5(ctx context.Context) { // 4-1 示例函数
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() { // 4-2 WithValue函数
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50) // 设置一个50毫秒的超时
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")          // 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	wg.Add(1)
	go worker5(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
