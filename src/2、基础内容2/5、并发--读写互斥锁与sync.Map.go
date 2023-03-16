package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var (
	//x       int64
	//wg      sync.WaitGroup
	mutex   sync.Mutex
	rwMutex sync.RWMutex
)

func writeWithLock() { // writeWithLock 使用互斥锁的写操作
	mutex.Lock() // 加互斥锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	mutex.Unlock()                    // 解互斥锁
	wg.Done()
}

func readWithLock() { // readWithLock 使用互斥锁的读操作
	mutex.Lock()                 // 加互斥锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	mutex.Unlock()               // 释放互斥锁
	wg.Done()
}

func writeWithRWLock() { // writeWithLock 使用读写互斥锁的写操作
	rwMutex.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwMutex.Unlock()                  // 释放写锁
	wg.Done()
}

func readWithRWLock() { // readWithRWLock 使用读写互斥锁的读操作
	rwMutex.RLock()              // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwMutex.RUnlock()            // 释放读锁
	wg.Done()
}

func do(wf, rf func(), wc, rc int) {
	start := time.Now()
	for i := 0; i < wc; i++ { // wc个并发写操作
		wg.Add(1)
		go wf()
	}

	for i := 0; i < rc; i++ { //  rc个并发读操作
		wg.Add(1)
		go rf()
	}

	wg.Wait()
	cost := time.Since(start)
	fmt.Printf("x:%v cost:%v\n", x, cost)

}
func main() {
	// 1-1使用互斥锁，10并发写，1000并发读
	do(writeWithLock, readWithLock, 10, 1000) // x:10 cost:1.466500951s

	// 1-2使用读写互斥锁，10并发写，1000并发读
	do(writeWithRWLock, readWithRWLock, 10, 1000) // x:10 cost:117.207592ms

	//wg := sync.WaitGroup{} //2-1 内置的 map 不是并发安全的 示例
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			//m.Lock() //加上锁
			set(key, n)
			//m.Unlock() //解开锁
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

var m1 = make(map[string]int)

func get(key string) int {
	return m1[key]
}

func set(key string, value int) {
	m1[key] = value
}

var m2 = sync.Map{} // 并发安全的map

func SyncMap() { //2-2 sync.Map 并发读写示例
	//wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ { // 对m执行20个并发的读写操作
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)         // 存储key-value
			value, _ := m2.Load(key) // 根据key取值
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
