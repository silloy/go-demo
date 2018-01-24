package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	var f = Adder1()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}

func Adder1() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
