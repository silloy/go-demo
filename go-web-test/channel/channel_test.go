package channel

import (
	"testing"
	"fmt"
)

func TestChan(t *testing.T) {
	run()

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	//v, ok := <-ch
	for i := range c {
		fmt.Println(i)
	}
}
