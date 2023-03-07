package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var number int = 4

func Increment() {
	time.Sleep(1 * time.Second)
	mutex.Lock()
	number--
	mutex.Unlock()
	wg.Done()
}

func Decrement() {
	mutex.Lock()
	number++
	wg.Done()
	mutex.Unlock()
}
func DecrementWithPanic1() {
	defer func() {
		if err := recover(); err != nil {
			// TODO create ticket here
			wg.Done()
			wg.Add(1)
			DecrementWithPanic1()
			fmt.Println("panic")
		}
	}()

	// wg.Done()
	panic("panic")
	// number++
	// wg.Done()
}

func DecrementWithPanic() {
	defer wg.Done()
	time.Sleep(10 * time.Second)

	number++
	// wg.Done()
}
func main() {
	// wg = sync.WaitGroup{}
	// // mutex = sync.Mutex{}
	// // wg.Add(2000)
	// // for i := 0; i < 1000; i++ {
	// // 	go Increment()
	// // 	go Decrement()
	// // }
	// // wg.Wait()
	// // fmt.Println(number)

	// wg.Add(2)
	// go DecrementWithPanic()
	// go DecrementWithPanic1()

	// wg.Wait()
	// var a interface{}
	// a = 2
	// if b, ok := a.(int); ok {
	// 	fmt.Println(b)
	// }
	// append()
	// fmt.Println(b)

	var a rune = 3332

	fmt.Printf("%c", a) 
	c := "Akshay"
	b := []uint8(c)
	fmt.Println(b)
	// b := byte(97)
	// fmt.Printf("%T", a)
	// fmt.Println()
	// fmt.Printf("%T,%c", b, b)
	// var c byte = 232
	// fmt.Printf("%T", c)
}

// how to recover

// defer func(){
// 	if err := recover(); err != nil {

// 	}
// }()
