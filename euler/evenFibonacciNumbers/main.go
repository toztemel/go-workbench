package main

import "fmt"

func main() {
	small := 1
	big := 2

	max := 4000000
	var sum int

	for big < max {
    if big%2 == 0 {
			sum += big
		}
		next := small + big
		small = big
		big = next
  	fmt.Println(small, " ", big)
	}
	fmt.Println("sum is:", sum)
}
