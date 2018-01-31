package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
			time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()



	/**
	  不要将关闭一个打开文件的操作和关闭一个channel操作混淆。
	  对于每个打开的文件，都需要在不使用的使用调用对应的Close方法来关闭文件。
	 */
	go func() {
		//for {
		//	x, ok := <- naturals
		//	if !ok {
		//		break  // channel was closed and drained
		//	}
		//	squares <- x*x
		//}
		for x := range naturals {
			squares <- x*x
		}
		close(squares)
	}()

	for {
		fmt.Println(<-squares)
	}
}
