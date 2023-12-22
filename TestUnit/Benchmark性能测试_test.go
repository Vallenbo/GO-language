package main

import (
	"testing"
	"time"
)

// 使用性能比较函数做测试的时候一个容易犯的错误就是把b.N作为输入的大小
//func BenchmarkFibWrong(b *testing.B) { // 错误示范1
//	for n := 0; n < b.N; n++ {
//		Fib(n)
//	}
//}

//func BenchmarkFibWrong2(b *testing.B) { // 错误示范2
//	Fib(b.N)
//}

func Fib(n int) int { // Fib 是一个计算第n个斐波那契数的函数
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func BenchmarkSplit(b *testing.B) { //4-1 性能基准测试示例
	for i := 0; i < b.N; i++ { // b.N是标准
		Split("沙河有沙又有河", "沙")
	}
}

func benchmarkFib(b *testing.B, n int) { //4-2 编写的性能比较函数
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B)  { benchmarkFib(b, 1) }
func BenchmarkFib2(b *testing.B)  { benchmarkFib(b, 2) }
func BenchmarkFib3(b *testing.B)  { benchmarkFib(b, 3) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(b, 10) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(b, 20) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(b, 40) }

func Benchmark3Split(b *testing.B) { //重置时间
	time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
	b.ResetTimer()              // 重置计时器
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

func BenchmarkSplitParallel(b *testing.B) { //5-1 并行测试示例
	b.RunParallel(func(pb *testing.PB) { // b.SetParallelism(1) // 设置使用的CPU数
		for pb.Next() {
			Split("沙河有沙又有河", "沙")
		}
	})
}