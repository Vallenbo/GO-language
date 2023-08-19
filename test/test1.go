package main

import (
	"fmt"
)

func P_type() { //打印类型
	a := "qwer"
	for i := 0; i < len(a); i++ {
		fmt.Printf("%T\n", a[i])
	}
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Printf("%v", len(strs))
}
