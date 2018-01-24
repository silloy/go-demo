package main

import (
	"time"
	"log"
	"fmt"
	"runtime"
	"os"
)

func main() {
	defer printStack()
	bigSlowOperation()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

// 不要忘记defer语句后的圆括号，否则本该在进入时执行的操作会在退出时执行，
// 而本该在退出时执行的，永远不会被执行
func bigSlowOperation() {
	defer trace("bigSlowOperation")() //  don't forget the extra parentheses
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow

}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}
