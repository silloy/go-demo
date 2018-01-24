package main

import (
	"os"
	"strconv"
	"fmt"
	"gopl/ch2/tempconv"
)

func main()  {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		f := tempconv.Fahreheit(t)
		c := tempconv.Celsius(t)
		fmt.Println(f)
		fmt.Println(tempconv.FToC(f))
		fmt.Println(c)
		fmt.Println(tempconv.CToF(c))
		fmt.Printf("%g = %s, %s = %g\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}