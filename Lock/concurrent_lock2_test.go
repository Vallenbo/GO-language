package engine

import (
	"fmt"
	"testing"
	"time"
)

/**
* @Description
* @Author aiJiang <mr.wan16@petalmail.com>
* @Date 2024/10/9 15:49
**/

func Test_locks(t *testing.T) {
	manager := NewLockResourceManager(time.Second)
	go manager.deadlockChecker.checkForDeadlock()

	manager.AcquireLock("testResource", time.Second, func() {
		fmt.Printf("Lock on resource timed out.\n")
	})
	manager.ReleaseLock("testResource")
}
