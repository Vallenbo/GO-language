package main

import (
	"fmt"
	"sync"
)

func CondCase() {
	list := make([]int, 0)
	cond := sync.NewCond(&sync.Mutex{}) //返回一个新的互斥锁
	go readList(&list, cond)
	go readList(&list, cond)
	go readList(&list, cond)
	initList(&list, cond)
}

func initList(list *[]int, c *sync.Cond) { //主叫方，可以持有锁，也可以不持有锁
	c.L.Lock()
	defer c.L.Unlock()
	for i := 0; i < 10; i++ {
		*list = append(*list, i)
	}
	c.Broadcast() //唤醒所有等待协程
	//c.Signal() //信号唤醒一个等待 c 的 goroutine，如果有的话。
}

func readList(list *[]int, c *sync.Cond) { //被叫方，必须持锁
	c.L.Lock()
	defer c.L.Unlock()
	for len(*list) == 0 {
		print("readList wait !!")
		c.Wait()
	}
	fmt.Println("list 数据为：", *list)
}

func main() {
	CondQueueCase()
}

func (q *Queue) GetMany(n int) []int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	for len(q.list) < n {
		q.cond.Wait()
	}
	list := q.list[:n]
	q.list = q.list[:n]
	return list
}

func CondQueueCase() {
	q := NewQueue()
	var wg sync.WaitGroup
	for n := 1; n <= 10; n++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			list := q.GetMany(n)
			fmt.Printf("%d：%d\n", n, list)
		}(n)
	}
}

type Queue struct {
	list []int
	cond *sync.Cond
}

func NewQueue() *Queue {
	return &Queue{
		list: []int{},
		cond: sync.NewCond(&sync.Mutex{}),
	}
}