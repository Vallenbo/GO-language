package main

import "fmt"

func intSum(x int, y int) int { //返回参数
	return x + y
}

func sayHello() {
	fmt.Println("Hello 沙河") //没有返回参数
}

func intSum2(x ...int) int { //切片类型的参数
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func f9(x func() int) { //传入函数作为参数，返回值为int
}

func f7(x func() int) (t func(int, int) int) { //返回值为带参函数，有返回值函数
	return t
}

func main() {
	sayHello()
	ret := intSum(10, 20)
	fmt.Println(ret)

	//defer语句试用代码
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f6()) //6
}

func f1() int { //defer语句试用代码
	x := 5
	defer func() {
		x++ //最后执行，且修改的是x
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x) //x当成参数传进函数
	return 5
}

func f6() (x int) {
	defer func(x *int) {
		*x++
	}(&x) //指针x当成参数传进函数
	return 5 //
}
