package io

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test_mathLibrary(t *testing.T) {
	rand.NewSource(time.Now().UnixNano()) //使用随机数种子 //使用时间纳秒级，保证每次不一样
	for i := 0; i < 5; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}
