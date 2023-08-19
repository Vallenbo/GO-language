package main

func climbStairs(n int) int { //动态规划
	if n < 3 {
		return n
	}

	num := [3]int{1, 2}
	for i := 3; i <= n; i++ {
		num[2] = num[0] + num[1]
		num[0] = num[1]
		num[1] = num[2]
	}

	return num[2]
}

func main() {

}