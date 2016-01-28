package main

import "fmt"

/*
Write a program that prints the value of this expression:
 (true && false) || (false && true) || !(false && false)
*/

func main() {
	value := (true && true) || (false && true) || !(false && false)
	fmt.Println (value)
}
