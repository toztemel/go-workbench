package main
/*
Write a function which takes an integer and returns two values:
the integer divided by 2
whether or not the integer is even (true, false)
 For example
half(1) should return (0, false)
half(2) should return (1, true).
*/
import "fmt"

func main() {

	half := func(n int) (int, bool) {
		return n/2, n%2==0
	}

  fmt.Println(half(5))
}
