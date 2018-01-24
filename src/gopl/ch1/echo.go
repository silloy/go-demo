package main

import (
	"os"
	"fmt"
	"strings"
)

func main()  {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args)
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	echo2()
	echo3()
}

func echo2()  {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3() {
	fmt.Println(strings.Join(os.Args[1: ], " "))
}
