package main

// go:generate ls -alh
//go:generate stringer -linecomment -type ErrCode ./src

import (
	"demo/src"
	"fmt"
)

func main() {
	fmt.Println(src.ERR_CODE_CONN_REFUSE)
}
