package main

import "fmt"


func Add_and_mulitply(x , y int) (int, int){
	addition := Add(x,y)
	multipation := Multiply(x,y)
	return addition, multipation
}

func main(){
	p, q := Add_and_mulitply(4,5)
	fmt.Println(p,q)
}