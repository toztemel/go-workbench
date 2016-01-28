package main

import "fmt"

func main() {
  largestPrime := findfactor(13195)
  fmt.Println(largestPrime)
  largestPrime = findfactor(600851475143)
  fmt.Println(largestPrime)
}

func findfactor(number int) (int) {
  divisor := 2

  for number > 0 && divisor!=number {
      if (number % divisor ==0) {
        number = number/divisor
      } else {
        divisor+=1
      }
  }
  return number;
}
