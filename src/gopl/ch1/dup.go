package main

import (
	"bufio"
	"os"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	counts := make(map[string]int)
	// bufio: process input and output
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()] ++
	}

	/**
	%d          十进制整数
	%x, %o, %b  十六进制，八进制，二进制整数。
	%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
	%t          布尔：true或false
	%c          字符（rune） (Unicode码点)
	%s          字符串
	%q          带双引号的字符串"abc"或带单引号的字符'c'
	%v          变量的自然形式（natural format）
	%T          变量的类型
	%%          字面上的百分号标志（无操作数
	 */

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s%n", n, line)
		}
	}
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// os.Open函数返回两个值。第一个值是被打开的文件(*os.File），其后被Scanner读取
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

func dup3() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
