package main

import "fmt"

func main() {
	var sum int

	for i := 999; i >= 0; i-- {
		if i%3 == 0 || i%5 == 0{
			sum += i
		}
	}
	fmt.Println(sum)
}
