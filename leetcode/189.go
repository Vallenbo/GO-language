package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%v\n", nums)
	rotate(nums, 3)
	fmt.Printf("%v\n", nums)
}

func rotate(nums []int, k int) {
	lens := len(nums)
	nums1 := make([]int, lens)
	copy(nums1, nums)
	for i, j := 0, 0; j < lens; j++ {
		if i+k >= lens {
			nums[i+k-lens] = nums1[i]
			i = i + k - lens
		} else {
			nums[i+k] = nums1[i]
			i = i + k
		}
	}
}
