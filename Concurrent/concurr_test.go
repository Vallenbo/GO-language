package Concurrent

import (
	"fmt"
	"testing"
)

func Test_concur(t *testing.T) {
	wg.Add(2)

	go addOperate()
	go addOperate()

	wg.Wait()
	fmt.Println(x)
}

func addOperate() { // add 对全局变量x执行5000次加1操作
	for i := 0; i < 5000; i++ {
		m.Lock() // 修改x前加锁
		x = x + 1
		m.Unlock() // 改完解锁
	}
	wg.Done()
}
