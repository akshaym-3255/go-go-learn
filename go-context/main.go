package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {

	fmt.Println("do something called")
	ctx, cancelCtx := context.WithCancel(ctx)
	ch := make(chan int)
	go doAnother(ctx, ch)

	for i := 0; i < 4; i++ {
		ch <- i
	}
	cancelCtx()
	time.Sleep(10 * time.Second)
	fmt.Println("do something done")
}

func doAnother(ctx context.Context, ch <-chan int) {
	fmt.Println("do another called ")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		case num := <-ch:
			fmt.Println(num)

		}
	}
}

type akshay struct {
	x int "akshay x"
	y int "akshay y"
}

func main() {
	// ak := akshay{
	// 	1, 2,
	// }

	// akt := reflect.TypeOf(ak)
	// x_tag, _ := akt.FieldByName("x")
	// fmt.Println(x_tag.Tag)
	ctx := context.Background()
	handler(ctx)
}

func handler(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(2*time.Second))
	defer cancel()

	// doSomething(ctx)
	Operation(ctx)

	err := databaseOprationThatwilltakeTime(ctx)
	if err != nil {
		fmt.Println("canceled function")
	}
}

func Operation(ctx context.Context) {
	time.Sleep(3 * time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("cancel operation after 2 second as database taking time")
		return
	default:
		fmt.Println("akshay")
	}
}
func databaseOprationThatwilltakeTime(ctx context.Context) error {

	time.Sleep(2 * time.Second)
	return errors.New("trest")

}
