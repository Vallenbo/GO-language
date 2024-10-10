package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	//println(max(1, 2))
	//print(min(1, 2))
	//println(canCompleteCircuit([]int{2, 3, 4}, []int{3, 4, 3}))
	//println(1000 % 1000)
	//println(1120 / 1000)

	// str := []string{"qwe", "ewq", "qww"}
	//s := "s"
	//for _, v := range s {
	//	println(string(v))
	//}

	//Tes_string()
	//print(isHappy(19))
	//print(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))

	print(minJumps([]int{2, 2, 1, 0, 1}))
}

func minJumps(nums []int) int {
	// 维护能够跳到的最远索引
	farthest := 0
	// 当前跳跃的结束索引
	currentEnd := 0
	jumps := 0

	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == currentEnd {
			currentEnd = farthest
			jumps++
		}
		if currentEnd >= len(nums)-1 {
			break
		}
	}
	return jumps
}

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	i := 0
	length := len(intervals)
	// 添加所有在 newInterval 之前的区间
	for i < length && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// 合并所有与 newInterval 重叠的区间
	for i < length && intervals[i][0] <= newInterval[1] {
		newInterval[0] = int(math.Min(float64(newInterval[0]), float64(intervals[i][0])))
		newInterval[1] = int(math.Max(float64(newInterval[1]), float64(intervals[i][1])))
		i++
	}
	result = append(result, newInterval)

	// 添加所有在 newInterval 之后的区间
	result = append(result, intervals[i:]...)

	return result
}

func minSubArrayLen(target int, nums []int) int {
	length := len(nums)
	front, rear, n := 0, 0, length+1
	sum := nums[front]

	for rear < length {
		if sum < target {
			rear++
			if rear < length {
				sum += nums[rear]
			}
		} else {
			if rear-front+1 < n {
				n = rear - front + 1
			}
			sum -= nums[front]
			front++
		}
	}
	if n > length {
		return 0
	}
	return n
}

func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	var gasnum, point, step int
	for cur := 0; cur < length; cur++ {
		gasnum = 0
		step = 0
		point = cur
		for {
			if step == length {
				return cur
			}
			gasnum += gas[point] - cost[point]
			if gasnum < 0 { // 没油
				break
			}
			point = (point + 1) % length
			step++
		}
	}
	return -1
}

func intToRoman(num int) string {
	romanStr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanNum := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	length := len(romanNum)
	var res strings.Builder
	var x int
	for i := 0; i < length; i++ {
		x = num / romanNum[i]
		if x > 0 {
			res.WriteString(strings.Repeat(string(romanStr[i]), x))
			num -= x * romanNum[i]
		}
	}
	return res.String()
}

func sortString(s string) string {
	// 将字符串转换为字符切片
	runes := []rune(s)
	// 对字符切片进行排序
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// 将排序后的字符切片转换回字符串
	return string(runes)
}

func Tes_string() {
	str := "asdfga"
	for _, v := range str {
		print(string(v))
	}
}

func test1() {
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

func isHappy(n int) bool {
	exist := map[int]bool{}
	var res int
	for {
		for n > 0 {
			res += (n % 10) * (n % 10)
			n /= 10
		}
		if res == 1 {
			return true
		} else if exist[res] {
			return false
		} else {
			exist[res] = true
		}
		n = res
		res = 0
	}
}
