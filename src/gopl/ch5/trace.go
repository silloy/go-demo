package main

import (
	"time"
	"log"
)

func main()  {
	bigSlowOperation()
}
// 不要忘记defer语句后的圆括号，否则本该在进入时执行的操作会在退出时执行，
// 而本该在退出时执行的，永远不会被执行
func bigSlowOperation()  {
	defer trace("bigSlowOperation")()  //  don't forget the extra parentheses
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