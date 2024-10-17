package slice

import (
	"fmt"
	"testing"
)

func Test_sliceAdd(t *testing.T) {
	s := []int{1, 2, 3}
	s1 := []int{1, 2, 3}
	s1 = append(s1, s...)

	fmt.Println(s1)
}
