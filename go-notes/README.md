# GO BASICS
___
## Install go 
[go documentation](https://go.dev/doc/install)

___
## Hello World
___
```go
package main

import "fmt"

func main() {
	fmt.Println("HELLO WORLD")
}
```

## Variables
___
```go
package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128
	var name = "akshay"
	var age int32 = 123;

	//if we declare a variable with const we cannot reassign it
	const isCool  = true
	// if we declare a variable in go we must need to use it else compiler will throw a error

	// shorthand for var declaration
	// we cannot use shorthand outside function in go there we need to use var 
	surname := "mohite"

	fmt.Println(name, age)

	//to get the type of variable we use %T verb
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", isCool)
	fmt.Printf("%T\n", 	surname)

}
```

___
## Packages

``` bash
packages
│   ├── main.go
│   └── str_util
│       ├── reverse.go
```

 reverse.go
```go
package str_util

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
```

 main.go

```go
package main

import (
	"fmt"
	"math"
	"/packages/str_util"
)

func main() {
	fmt.Println("Hello world from package")
	fmt.Println(math.Sqrt(25))
	fmt.Println(str_util.Reverse("hello"))
}

```
___

## Functions
```go
package main

import "fmt"

func Greeting(name string) string{
	return "hello " + name;
}

func GetSum(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(Greeting("akshay"))
	fmt.Println(GetSum(4,5))
}

```
___

## Arrays and Slices

```go
package main

import "fmt"

func main() {
	//defining array 
	var fruit_array [2]string

	fruit_array[0] = "apple"
	fruit_array[1] = "mango"
	fmt.Println(fruit_array[0])
	fmt.Println(fruit_array[:2])

	//declare and assign at same time
	 flower_array := []string{"rose","lotus"}
	fmt.Println(flower_array[0])

	//length of array
	fmt.Println(len(flower_array))
	
	//slice an aray
	fmt.Println(flower_array[1:2])

	//partial intilzation 0 by default value empty string in case of string 
	
	teams_array := [5]string{"india","england"}

	fmt.Printf("%T",teams_array[3])
	fmt.Println(len(teams_array))

	//range is iterator over multiple data types
	for k := range flower_array{
		fmt.Println(k)
	}

	//slices are just like arrays we dont need to provide length while defining

	// sliceExample make([]string)

}
```

___

## Conditionals

```go
package main

import (
	"fmt"
)

func main() {
	x := 5
	y := 6

	if x < y {
		fmt.Println("x is less than y")
	} else if x == y {
		fmt.Println("y is equal to x")
	} else {
		fmt.Println("y is less than x")
	}

	//switch case
	switch x {
	case 5:
		fmt.Println("value of x is 5")
	case 6:
		fmt.Println("value of x is 6")
	default:
		fmt.Println("value of x other than 5 and 6")
	}
}
```
___


## Loops

```
package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}


	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Print("fizbuzz\n")
		} else if i%3 == 0 {
			fmt.Print("fiz\n")
		} else {
			fmt.Print("buzz\n")
		}
	}
}
```
___

## Maps

```go
package main

import (
	"fmt"
)

func main() {
	//declaration of map
	emails := make(map[string]string)
	emails["aksahy"] = "akshay@gmail.com"
	emails["bob"] = "bob@gmail.com"
	fmt.Println(emails)

	//iterating into email
	for k := range emails {
		fmt.Println(k,emails[k])
	}

	email2 := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}

	email2["akshay"] = "akshay@gmail.com"

	//delete from map
	delete(email2,"Bob")

	fmt.Println(email2)
}
```
___

## Range

```go
package main

import (
	"fmt"
)

func main() {
	// define a array
	numbers := []int{1,2,3,4,5}

	// iterate through index and element
	for i, number := range numbers{
		fmt.Println(i,number)
	}

	// iteratr thorugh index only
	for i := range numbers{
		fmt.Println(numbers[i])
	}

	//map defination
	contacts := map[string]int{"akshay":12345,"bob":13315315}

	// iteratin through key and value
	for k,v := range contacts {
		fmt.Println(k,v)
	}

	// iterating over only key 
	for k := range contacts{
		fmt.Println(k)
	}
	
}
```
___

## Pointers

```go
package main

import (
	"fmt"
)

func paas_by_value(a int) {
	a = 10
}

func pass_by_reference(a *int){
	*a = 15
}

func main() {
	a := 4
	b := &a
	fmt.Println(a,b)
	fmt.Printf("%T\n",b)


	fmt.Println(*b)
	fmt.Println(*&a)

	//change the value of a
	*b = 10
	fmt.Println(a)

	paas_by_value((a))
	fmt.Println(a)

	pass_by_reference(&a)
	fmt.Println(a)

}
```
___

## Closures

```go
package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}
```
___

## Struct

```go 

```

## go MOD file 

