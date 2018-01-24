package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

const (
	Avogadro = 6.02214129e23  // 阿伏伽德罗常数
	Planck   = 6.62606957e-34 // 普朗克常数
)

func main() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"

	char()
	countrune()
}


func char() {
	sa := "😀"
	fmt.Println(len(sa))

	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"


	n := 0
	for i, r := range "Hello, 世界" {
		n++
		fmt.Printf("%d\t%q\t%d\t%d\n", i, r, r, n)
	}
}

func countrune()  {
	s := "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r)  // "[30d7 30ed 30b0 30e9 30e0]"
	fmt.Println(string(r))
	fmt.Println(string(65))     // "A", not "65"
	fmt.Println(string(0x4eac)) // "京"
	fmt.Println(string(1234567)) // "�"
}

