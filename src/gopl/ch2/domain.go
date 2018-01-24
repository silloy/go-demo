package main

import (
	"fmt"
	"os"
)

func main() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
			fmt.Println()
		}
	}
	trans()
}

func trans()  {
	x := "hello"
	fmt.Println(os.Getwd())
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}
}