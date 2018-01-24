package main

import (
	"fmt"
)

func main() {
	fmt.Println(gcd(5, 7))
	fmt.Println(fib(10))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
