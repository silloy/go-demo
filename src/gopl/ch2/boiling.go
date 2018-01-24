package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	boiling()
	pointer()

	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
}

func boiling() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g °C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g °C\n", boilingF, fToC(boilingF))
}

func fToC(f float64) interface{} {
	return (f - 32) * 5 / 9
}

func pointer()  {
	x := 1
	p := &x
	fmt.Println(p)
	fmt.Println(*(&x))
	*p = 2
	fmt.Println(x)

	var x1, y int
	fmt.Println(&x1 == &x1, &x1 == &y, &x1 == nil) // "true false false"
}


/**
 在Go语言中，返回函数中局部变量的地址也是安全的。例如下面的代码，
 调用f函数时创建局部变量v，在局部变量地址被返回之后依然有效，因为指针p依然引用这个变量。
 */

var p = f()

func f() *int {
	v := 1
	fmt.Println(&v)
	return &v
}


func incr(p *int) int {
	*p++ // 非常重要：只是增加p指向的变量的值，并不改变p指针！！！
	return *p
}



