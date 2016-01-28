package main

import "fmt"

func main() {
	fmt.Println("Enter your name and surname")
  input := ""
  surname := ""
  fmt.Scan(&input, &surname)
	fmt.Println("Hello", input, surname)
}
