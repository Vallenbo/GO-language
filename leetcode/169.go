package main

import "fmt"

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	i := 0
	num := float32(len(nums)) / 2
	for ; i < len(nums); i++ {
		mark := float32(1)
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				mark++
			}
			if mark > num {
				return nums[i]
			}
		}
	}
	return nums[i]
}

func main() {
	num := []int{2, 2, 3, 3, 3, 3, 2}
	fmt.Println(majorityElement(num))
}