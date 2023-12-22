package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	IntJob := make(chan int, 10)       //create work pool
	StringJob := make(chan string, 10) //create work pool
	for i := 1; i <= 10; i++ {         //work pool add water
		IntJob <- i
		StringJob <- fmt.Sprint("str + ", strconv.Itoa(i))
	}

	wg2 := sync.WaitGroup{}
	wg2.Add(10)

}
