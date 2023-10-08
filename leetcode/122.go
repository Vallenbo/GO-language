package main

import "fmt"

func main() {
	nums := []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("%v\n", nums)
	profit := maxProfit(nums)
	fmt.Printf("%v\n", profit)
}

func maxProfit(prices []int) int {
	profit := 0
	rear := 1
	front := 0
	for rear < len(prices) {
		if prices[front] < prices[rear] {
			profit = prices[rear] - prices[front] + profit
			front = rear
			rear++
		} else {
			front++
			rear++
		}
	}
	return profit
}
