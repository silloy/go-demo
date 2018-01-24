package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing neweline")
var sep = flag.String("s", "", "separtor")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println("-----")
	}
	newF()
}

func newF() {
	p := new(int)   // p, *int 类型, 指向匿名的 int 变量
	fmt.Println(*p) // "0"
	*p = 2          // 设置 int 匿名变量的值为 2
	fmt.Println(*p) // "2"
}

// 由于new被定义为int类型的变量名，因此在delta函数内部是无法使用内置的new函数的。
func delta(old, new int) int {
	return new - old
}



var global *int

// f函数里的x变量必须在堆上分配，因为它在函数退出后依然可以通过包一级的global变量找到
func ff() {
	var x int
	x = 1
	global = &x
}

// *y并没有从函数g中逃逸，编译器可以选择在栈上分配*y的存储空间
// （译注：也可以选择在堆上分配，然后由Go语言的GC回收这个变量的内存空间）
func g() {
	y := new(int)
	*y = 1
}
