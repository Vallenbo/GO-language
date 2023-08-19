package main

import (
	"fmt"
)

func main() {
	s := [3]string{"s", "a", "b"}
	var s1 [2]string
	fmt.Printf("%v,%v\n", s, s1)
}

func test() {
	a := make([]map[int]string, 2, 2) //初始化切片
	b := make(map[bool]int)           //初始化map
	fmt.Printf("%v,%v\n", a, b)
	//a := make(map[int]bool)
	//b := make([]map[int]string)

}
