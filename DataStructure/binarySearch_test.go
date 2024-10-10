package DataStructure

import (
	"fmt"
	"testing"
)

func binarySearch(arr []int, target int) int { //二分查找使时间复杂度：log（N）
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func Test_binarySearch(*testing.T) {
	arr := []int{1, 3, 5, 7, 9}

	fmt.Println(binarySearch(arr, 3))  // Output: 1
	fmt.Println(binarySearch(arr, 10)) // Output: -1
}
