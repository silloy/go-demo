package main

import (
	"fmt"
	"gopl/ch2/tempconv"
	"flag"
	"io"
	"os"
	"bytes"
)

type celsiusFlag struct {
	tempconv.Celsius
}

// flag.value interface
func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	// fmt.Sscanf函数从输入s中解析一个浮点数（value）和一个字符串（unit）
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = tempconv.Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = tempconv.FToC(tempconv.Fahreheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C".
func CelsiusFlag(name string, value tempconv.Celsius, usage string) *tempconv.Celsius {
	f := celsiusFlag{value}
	/**
	  Celsius字段是一个会通过Set方法在标记处理的过程中更新的变量。调用Var方法将标记加入应用的命令行标记集合中，
	有异常复杂命令行接口的全局变量flag.CommandLine.Programs可能有几个这个类型的变量。调
	用Var方法将一个*celsiusFlag参数赋值给一个flag.Value参数,导致编译器去检查*celsiusFlag是否有必须的方法
	 */
	flag.CommandLine.Var(&f, name, usage)
	os.Stdout.Write([]byte("hello")) // "hello"
	var w io.Writer
	fmt.Printf("%T\n", w) // "<nil>"
	w = os.Stdout
	fmt.Printf("%T\n", w) // "*os.File"
	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) // "*bytes.Buffer"
	return &f.Celsius
}
