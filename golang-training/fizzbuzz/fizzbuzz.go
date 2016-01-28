package main

import "fmt"

func main() {

	for i := 0; i < 100; i++ {
    if i%15 == 0 {
			fmt.Println("FizzBazz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Bazz")
		} else {
			fmt.Println(i)
		}
	}

}
