package main

import (
	"fmt"
	"sync"
	"time"
)

func Async(semaphore chan struct{}, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	semaphore <- struct{}{}
	fmt.Println(i)
	time.Sleep(2 * time.Second)
	<-semaphore
}

func main() {
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 2)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Async(semaphore, &wg, i)
	}
	wg.Wait()
}
