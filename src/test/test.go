package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"John", "Mary", "Peter", "Bob", "Tom", "Jane"}
	coinMap := make(map[rune]int) // 用于存储每个字母对应的金币数量
	// 定义每个字母的金币数量
	coinMap['a'] = 1
	coinMap['e'] = 1
	coinMap['i'] = 2
	coinMap['o'] = 3
	coinMap['u'] = 5
	for _, name := range names {
		coins := 0
		for _, char := range strings.ToLower(name) { // 将名字转换为小写字母，以便不区分大小写
			if coin, ok := coinMap[char]; ok {
				coins += coin
			}
		}
		fmt.Printf("%s: %d\n", name, coins)
	}
}
