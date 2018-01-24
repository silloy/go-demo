package main

import (
	"fmt"
)

// #include <stdio.h>
// #include <stdlib.h>
import (
	"os"
	"runtime"
	"math/rand"
	"unicode/utf8"
	"strconv"
	"time"
)

func main() {
	fmt.Println("测试")

	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)

	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	var o = 3
	fmt.Println(i, j, k, l, &o)

	var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	fmt.Println(balance, len(balance))

	var f int = 10

	fmt.Printf("变量的地址: %x\n", &f)

	book := books{"title", "tao", 1}
	fmt.Println(book, book.author)

	s := []int{1, 2, 3}
	s1 := make([]int, 2)
	fmt.Println(s, s1)

	//nums := []int{2,3,4}
	kvs := map[string]string{"a": "apple"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
	fmt.Println(Divide(6, 2))
	fmt.Printf("%+v%#v%T", "adf", "Sd", "sdfdfd")
	fmt.Println("a");
	goos := os.Getenv("GOOS")
	fmt.Printf("The operating system is: %s\n", goos)
	fmt.Printf("The operating system is: %s\n", runtime.GOOS)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)

	//ss := "dfd";
	//cs := C.CString(ss)
	//defer C.free(unsafe.Pointer(cs))
	//C.fputs(cs, (*C.FILE)(C.stdout))
	n()
	m()
	n()

	a1, b1 := 10, 2
	c1 := a1 / b1 // panic: runtime error: integer divide by zero

	fmt.Println(c1)

	fmt.Println(rand.Intn(233434), "-------------")

	str1 := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("The number of bytes in string str1 is %d\n",len(str1))
	fmt.Printf("The number of characters in string str1 is %d\n",utf8.RuneCountInString(str1))
	str2 := "asSASA ddd dsjkdsjsこん dk"
	fmt.Printf("The number of bytes in string str2 is %d\n",len(str2))
	fmt.Printf("The number of characters in string str2 is %d",utf8.RuneCountInString(str2))


	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)

	fmt.Println(time.Now().UTC())

	fmt.Println(getX2AndX3_2(6))

	v := make([]int, 10, 50)
	fmt.Println(v)



}

func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	 //return x2, x3
	return
}


var a = "G"

func n() { fmt.Println(a) }

func m() {
	a := "O"
	fmt.Println(a)
}

type books struct {
	title   string
	author  string
	book_id int
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func Divide(a int, b int) (result int, error string) {
	if b == 0 {
		data := DivideError{
			dividee: a,
			divider: b,
		}
		error = data.Error()
		return
	} else {
		return a / b, ""
	}
}

// 实现 	`error` 接口
func (de *DivideError) Error() string {
	strFormat := `
	Cannot proceed, the divider is zero.
	dividee: %d
	divider: 0
	`
	return fmt.Sprintf(strFormat, de.dividee)
}

type DivideError struct {
	dividee int
	divider int
}

func init() {
	fmt.Println("init.............")
}
