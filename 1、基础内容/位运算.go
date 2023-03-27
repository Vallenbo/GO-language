package main

import "fmt"

func test0() { //1 - 依次取出二进制最右边的1  //依次取出二进制数字
	a := 148 // 10010100
	for a != 0 {
		rightOne := a & (-a)
		fmt.Println(rightOne)
		a ^= rightOne
	}
}

func oddFind() { //2- ^使用用例：数组中只有一种数出现了奇数次找到它
	arr1 := []int{4, 3, 7, 2, 7, 2, 3, 2, 7, 2, 3, 3, 4}
	eor := 0
	for i := 0; i < len(arr1); i++ {
		eor ^= arr1[i]
	}
	println("To present odd number is :", eor)
}

func TwoOddFind() { //2-1 ^使用用例：数组中有2种数出现了奇数次找到它
	arr2 := []int{4, 3, 1, 7, 2, 7, 2, 3, 2, 7, 2, 3, 3, 4}
	eor := 0
	for i := 0; i < len(arr2); i++ {
		eor ^= arr2[i]
	}
	rightOne := eor & (-eor)
	onlyOne := 0
	for i := 0; i < len(arr2); i++ {
		if (arr2[i] & rightOne) != 0 {
			onlyOne ^= arr2[i]
		}
	}
}

func main() {

	// 按位与 &
	a0 := 5                                  // 二进制表示为 101
	b0 := 3                                  // 二进制表示为 011
	c0 := a0 & b0                            // 二进制为 001，即十进制为 1
	fmt.Printf("%d & %d = %d\n", a0, b0, c0) //5 & 3 = 1

	// 按位或 |
	a0 = 5                                   // 二进制表示为 101
	b0 = 3                                   // 二进制表示为 011
	c0 = a0 | b0                             // 二进制为 111，即十进制为 7
	fmt.Printf("%d | %d = %d\n", a0, b0, c0) //5 | 3 = 7

	// 按位异或 ^
	a0 = 5                                   // 二进制表示为 101
	b0 = 3                                   // 二进制表示为 011
	c0 = a0 ^ b0                             // 二进制为 110，即十进制为 6
	fmt.Printf("%d ^ %d = %d\n", a0, b0, c0) //5 ^ 3 = 6

	// 左移 <<
	a0 = 5                               // 二进制表示为 101
	b0 = a0 << 1                         // 二进制为 1010，即十进制为 10
	fmt.Printf("%d << 1 = %d\n", a0, b0) //5 << 1 = 10

	// 右移 >>
	a0 = 5                               // 二进制表示为 101
	b0 = a0 >> 1                         // 二进制为 10，即十进制为 2
	fmt.Printf("%d >> 1 = %d\n", a0, b0) //5 >> 1 = 2
}
