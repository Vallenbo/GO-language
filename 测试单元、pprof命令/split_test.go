package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestSplit(t *testing.T) { // 1- 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
	got := Split("a:b:c", ":")         // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
		t.Errorf("expected:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}

func Test1Split(t *testing.T) { //2- 测试组
	type test struct { // 定义一个测试用例类型
		input string
		sep   string
		want  []string
	}

	tests := []test{ // 定义一个存储测试用例的切片
		{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		{input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		{input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}

	for _, tc := range tests { // 遍历切片，逐一执行测试用例
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("expected:%v, got:%v", tc.want, got)
		}
	}
}

func Test3Split(t *testing.T) { //3-1 子测试用例
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	for name, tc := range tests {
		got := Split(tc.input, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("name:%s expected:%#v, got:%#v", name, tc.want, got) // 将测试用例的name格式化输出
		}
	}
}

func Test4Split(t *testing.T) { //3-2 子测试用例 ，Go1.7+中新增了子测试，使用t.Run执行子测试
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"河有", "又有河"}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}

func BenchmarkSplit(b *testing.B) { //4-1 性能基准测试示例
	for i := 0; i < b.N; i++ { // b.N是标准
		Split("沙河有沙又有河", "沙")
	}
}

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

func TestMain(m *testing.M) { //6-1 estMain
	fmt.Println("write setup code here...") // 测试之前的做一些设置
	// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         // 执行测试
	fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
	os.Exit(retCode)                           // 退出测试
}

func setupTestCase(t *testing.T) func(t *testing.T) { //7-1  测试集的Setup与Teardown
	t.Log("如有需要在此执行:测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:测试之后的teardown")
	}
}

func setupSubTest(t *testing.T) func(t *testing.T) { // 子测试的Setup与Teardown
	t.Log("如有需要在此执行:子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行:子测试之后的teardown")
	}
}

func Test5Split(t *testing.T) {
	type test struct { // 定义test结构体
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{ // 测试用例使用map存储
		"simple":      {input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"wrong sep":   {input: "a:b:c", sep: ",", want: []string{"a:b:c"}},
		"more sep":    {input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"leading sep": {input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}
	teardownTestCase := setupTestCase(t) // 测试之前执行setup操作
	defer teardownTestCase(t)            // 测试之后执行testdoen操作

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) { // 使用t.Run()执行子测试
			teardownSubTest := setupSubTest(t) // 子测试之前执行setup操作
			defer teardownSubTest(t)           // 测试之后执行testdoen操作
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("expected:%#v, got:%#v", tc.want, got)
			}
		})
	}
}
