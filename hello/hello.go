package main

import  (
	"fmt"
	"github.com/toztemel/stringutil"
)

func main() {
	fmt.Printf("hello, world\n")
	fmt.Printf(stringutil.Reverse("\nhello, world!"))
	fmt.Printf("%c \n", 65)
	fmt.Printf("%d \n", 'A')
	fmt.Println([]byte("ABCabc"))
	fmt.Println([]byte("世界"))

	greet(world)
	greet("Tayfun")
}

func greet (name String) {
	fmt.Println("Hello", name)
}
