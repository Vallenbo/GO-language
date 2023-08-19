package main

import "fmt"

func isValid(s string) bool {
	stack := make([]string, 0) // 使用切片定义栈
	for _, val := range s {
		if isLeft(string(val)) { // 如果是左符号则进栈
			stack = append(stack, string(val))
		} else { // 此时为右符号，执行出栈
			if len(stack) == 0 {
				return false
			}
			if isFit(stack[len(stack)-1], string(val)) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func isLeft(char string) bool {
	switch char {
	case "(", "[", "{":
		return true
	}
	return false
} // 判断是否为左边

func isFit(char1 string, char2 string) bool {
	flag := false
	if char1 == "(" && char2 == ")" {
		flag = true
	} else if char1 == "[" && char2 == "]" {
		flag = true
	} else if char1 == "{" && char2 == "}" {
		flag = true
	}
	return flag
} // 判断是否合适

func main() {
	s := "()[]{}"
	fmt.Println(isValid(s))
}