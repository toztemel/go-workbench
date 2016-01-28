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
  b, even := half (2)
  fmt.Println(b, even)
}

func half(p int)(int, bool)  {
  return p/2, p%2==0
}
