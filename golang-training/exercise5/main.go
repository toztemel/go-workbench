package main

import "fmt"

/*
Write a program that prints the value of this expression:
 (true && false) || (false && true) || !(false && false)
*/

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(numbers ...int){
	fmt.Println(numbers)
}
