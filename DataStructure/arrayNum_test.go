package DataStructure

import (
	"fmt"
	"testing"
)

func Test_array(t *testing.T) { //预处理结构   //前缀和数组:使查询数组中某范围内累加和，有点快
	nums := []int{1, 2, 3, 4, 5}
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}

	fmt.Println("nums:", nums)
	fmt.Println("prefix sum:", prefixSum)
}
