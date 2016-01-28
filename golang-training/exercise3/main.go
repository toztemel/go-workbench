package main

import "fmt"

/*
Write a function with one variadic parameter
that finds the greatest number in a list of numbers.
*/

func main() {
	data := []int{1100, 4, 9, 11, 10, 1, 0, 92, 888}
	max := findMax(data...)
	fmt.Println(max)
}

func findMax(number ...int) int {
	max := 0
	for _, v := range number {
		if v > max {
			max = v
		}
	}
	return max
}
