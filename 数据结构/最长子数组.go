package main

func maxSubArrayLen(nums []int, k int) int {
	sumMap := make(map[int]int)
	sumMap[0] = -1
	maxLength, sum := 0, 0
	for i, num := range nums {
		sum += num
		if j, ok := sumMap[sum-k]; ok {
			maxLength = max(maxLength, i-j)
		}
		if _, ok := sumMap[sum]; !ok {
			sumMap[sum] = i
		}
	}
	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
