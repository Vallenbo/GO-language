package main

import "fmt"

type Cat struct{}

type animal interface {
	Say()
}

func (c Cat) Say() {
	fmt.Println("喵喵喵~")
}

type Dog1 struct{}

func (d Dog1) Say() {
	fmt.Println("汪汪汪~")
}

func action(a animal) {
	a.Say()
}

func main() {
	c := Cat{}

	c.Say()
	d := Dog1{}
	d.Say()
	action(c)

	assign(100)
	assign2(true)
}

func void_interface() { //2- 定义空接口类型的
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 4)
	m1["name"] = "李白"

	fmt.Printf("The type of m1:%T, value:%v\n", m1, m1)
}

func assign(a interface{}) { //3- 断言
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if ok {
		fmt.Println("类型断言成功")
	} else {
		fmt.Println("类型断言失败, 传进来的是： ", str)
	}
}

func assign2(a interface{}) { //3-2 对一个接口值有多个实际类型需要判断，推荐使用switch语句来实现
	fmt.Printf("%T\n", a)
	switch v := a.(type) { //类型断言，主要用switch
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupported type！")
	}
}
