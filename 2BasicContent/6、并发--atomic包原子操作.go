package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter interface {
	Inc()
	Load() int64
}

type CommonCounter struct { // 普通版
	counter int64
}

func (c CommonCounter) Inc() {
	c.counter++
}

func (c CommonCounter) Load() int64 {
	return c.counter
}

type MutexCounter struct { // 互斥锁版
	counter int64
	lock    sync.Mutex
}

func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

type AtomicCounter struct { // 原子操作版
	counter int64
}

func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1) //执行加操作，传入：值和加数
}

func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func test1(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

func main() {
	c1 := CommonCounter{} // 非并发安全
	test1(c1)
	c2 := MutexCounter{} // 使用互斥锁实现并发安全
	test1(&c2)
	c3 := AtomicCounter{} // 并发安全且比互斥锁效率更高
	test1(&c3)

	x = 200
	ok := atomic.CompareAndSwapInt64(&x, 100, 300) //原子操作，用100与x比较，为ture，则x=300
	fmt.Println(ok, x)
}
