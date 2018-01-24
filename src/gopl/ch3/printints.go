package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
	conv()
}

// intsToString is like fmt.Sprint(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func conv() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))              // "123 123"
	fmt.Println(strconv.FormatInt(int64(x), 2))  // "1111011"
	fmt.Sprintf("x=%b", x)                       // "x=1111011"
	fmt.Println(strconv.Atoi("123"))             // x is an int
	fmt.Println(strconv.ParseInt("123", 10, 64)) // base 10, up to 64 bits
}
