package main

import "fmt"

type Stack struct { // 定义栈结构体
	data []interface{} // 数据
}

func (stack *Stack) Push(data interface{}) { // 将元素压入栈
	stack.data = append(stack.data, data)
}

func (stack *Stack) Pop() interface{} { // 从栈中弹出元素
	if len(stack.data) == 0 {           // 如果栈为空，则返回nil
		return nil
	}

	data := stack.data[len(stack.data)-1]       // 取出栈顶元素
	stack.data = stack.data[:len(stack.data)-1] // 将栈顶元素从栈中删除
	return data
}

func (stack *Stack) Peek() interface{} { // 获取栈顶元素
	if len(stack.data) == 0 {            // 如果栈为空，则返回nil
		return nil
	}

	data := stack.data[len(stack.data)-1] // 取出栈顶元素
	return data
}

func (stack *Stack) IsEmpty() bool { // 判断栈是否为空
	return len(stack.data) == 0
}
func main() {
	stack := &Stack{}
	stack.Push(1) // 将元素压入栈
	stack.Push(2)
	stack.Push(3)

	fmt.Println(stack.Peek())    // 获取栈顶元素
	fmt.Println(stack.Pop())     // 从栈中弹出元素
	fmt.Println(stack.IsEmpty()) // 判断栈是否为空
}
