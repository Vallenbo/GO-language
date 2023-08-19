package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var a, b = 0, 1
	var nums1 []int
	for a < len(nums) {
		for b < len(nums) {
			if nums[a]+nums[b] == target {
				nums1 = append(nums1, a)
				nums1 = append(nums1, b)
				return nums1
			}
			b = b + 1
		}
		a = a + 1
		b = a + 1
	}
	return nums1
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	sum := twoSum(nums, target)
	fmt.Println(sum)
}