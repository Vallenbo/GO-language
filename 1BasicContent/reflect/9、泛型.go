package main

import "fmt"

func SimpleCase() {
	var a, b = 3, 4
	var c, d float64 = 5, 6
	fmt.Println("不使用泛型，数字比较：", getMaxNumInt(a, b))
	fmt.Println("不使用泛型，数字比较：", getMaxNumFloat(c, d))

	fmt.Println("使用泛型，数字比较：", getMaxNumInt(a, b))            //由编译器推断输入的类型
	fmt.Println("使用泛型，数字比较：", getMaxNumFloat[float64](c, d)) // 显示指定传入的类型
}

func main() {

}

func getMaxNumInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxNumFloat(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func getMaxNum[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type CusNumT interface {
	uint8 | int32 | float64 | ~int64 //~表示支持类型的衍生类型
}

type Myint64 int64 //衍生类型，是具有基础类型的int64的新类型，与int64是不同类型

type Myint32 = int32 //int32的别名，是同一类型

func getBuiltInComparable[T comparable](a, b T) bool { // comparable类型，只支持 == != 两个操作的任意类型
	if a == b {
		return true
	}
	return false
}

func BuiltCase() {
	var a, b = "abc", "efg"
	fmt.Println("内置comparable 泛型类型约束", getBuiltInComparable(a, b))
	var c, d float64 = 100, 100
	fmt.Println("内置comparable 泛型类型约束", getBuiltInComparable(c, d))
}