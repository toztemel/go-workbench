package main

import "fmt"

func main() {
	total := 0
	for i := 1000; i > 0; i-- {
		if i%3 == 0 {
			total += i
		} else if i%5 == 0 {
			total += i
		}
	}
	fmt.Println(total) //234168

}
