package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := []int{0, 1, 2, 4, 5, 7}
	fmt.Printf("%v\n", summaryRanges(s))
}

func summaryRanges(nums []int) []string {
	str := []string{}
	n := 1
	for i, rear := 0, 1; i < len(nums) && rear < len(nums); {
		if nums[i]+n == nums[rear] {
			rear++
			n++
		} else if i != rear {
			str = append(str, strconv.Itoa(nums[i])+"->"+strconv.Itoa(nums[rear-1]))
			i = rear
			if rear+1 < len(nums) {
				rear++
			}
			n = 1
		} else {
			str = append(str, strconv.Itoa(nums[i]))
			return str
		}
	}
	return str
}
