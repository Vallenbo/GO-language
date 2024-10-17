package main

import (
	"fmt"
	"sync"
)

func worker4(job <-chan int, wg *sync.WaitGroup, workerID int) {
	for j := range job {
		fmt.Printf("workerID:%d,jobID：%d\n", workerID, j)
		wg.Done()
	}
}

func main() {
	job := make(chan int, 10)  //create work pool
	for i := 1; i <= 10; i++ { //work pool add water
		job <- i
	}
	wg2 := sync.WaitGroup{}
	wg2.Add(10)

	for i := 1; i <= 3; i++ {
		go worker4(job, &wg2, i)
	}
	wg2.Wait()
	close(job)
}
