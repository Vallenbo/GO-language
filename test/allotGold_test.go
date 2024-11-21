package test

import (
	"fmt"
	"testing"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

// 定义每个字母对应的金币数量
var coinMap = map[rune]int{
	'e': 1,
	'E': 1,
	'i': 2,
	'I': 2,
	'o': 3,
	'O': 3,
	'u': 4,
	'U': 4,
}

// 计算每个用户分到多少金币，以及最后剩余多少金币
func dispatchCoin(names []string, coins int) (map[string]int, int) {
	coinCount := make(map[string]int) //硬币 姓名 ，计数变量
	for _, name := range names {      //循环姓名数组
		count := 0
		for _, char := range name {
			if coin, ok := coinMap[char]; ok {
				count += coin
			}
		}
		coinCount[name] = count //进行姓名和硬币分配赋值
		coins -= count          //总硬币数减少
	}
	return coinCount, coins
}

func Test_allotGold(*testing.T) {
	names := []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	coins := 50 // 初始金币数量
	coinCount, leftCoins := dispatchCoin(names, coins)
	fmt.Println("每个用户分到的金币数量：", coinCount)
	fmt.Println("剩余的金币数量：", leftCoins)
}
