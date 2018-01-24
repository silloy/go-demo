package main

import (
	"fmt"
	"time"
)

type Flags uint

const (
	_   = 1 << (10 * iota)
	KiB  // 1024
	MiB  // 1048576
	GiB  // 1073741824
	TiB  // 1099511627776             (exceeds 1 << 32)
	PiB  // 1125899906842624
	EiB  // 1152921504606846976
	ZiB  // 1180591620717411303424    (exceeds 1 << 64)
	YiB  // 1208925819614629174706176
)

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
	fmt.Println(YiB/ZiB)
	initArrray()
}

/**
	六种未明确类型的常量类型，分别是无类型的布尔型、无类型的整数、
	无类型的字符、无类型的浮点数、无类型的复数、无类型的字符串
 */

func constValue()  {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"
}

func initArrray()  {
	type Currency int

	const (
		USD Currency = iota // 美元
		EUR                 // 欧元
		GBP                 // 英镑
		RMB                 // 人民币
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}

	fmt.Println(RMB, symbol[RMB]) // "3 ￥"
}



