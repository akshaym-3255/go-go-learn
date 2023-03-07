package main

import (
	"fmt"
	"sync"
	"time"
)

var globalSlice []string

var wg *sync.WaitGroup

func add_values_to_slice(ch chan []string, x ...string) {
	local := []string{}
	local = append(local, x...)
	time.Sleep(3 * time.Second)
	ch <- local
}

func remove_values_from_slice(ch chan []string) {
	local := []string{"akshay"}
	local = append(local, "afaf")
	time.Sleep(3 * time.Second)
	ch <- local

}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("pnaicked and recovred")
		}
	}()
	wg = &sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan []string)
	ch2 := make(chan int)
	go add_values_to_slice(ch, "akshay", "shubham", "shree")
	// fmt.Println(<-ch)
	go remove_values_from_slice(ch)

	go func(val int) {
		for i := 0; i < 4; i++ {
			ch2 <- i
		}
		close(ch2)
	}(4)
	fmt.Println(<-ch)
	fmt.Println("asfafdfa")
	fmt.Println(<-ch)
	// fmt.Println(<-ch2)
	// panic("panicked")

	for c := range ch2 {
		fmt.Println(c)
	}
}
