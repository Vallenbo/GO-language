package main

import "fmt"

func tank(n int) int {
	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	return tank(n-1) + tank(n-2)
}

func main() {
	//n个台阶，一次可以走1个台阶，也可以走2个台阶，有多少种走法
	fmt.Println("Please enter the number of steps :")
	n := 0
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	print("the method is :", tank(n))

}
