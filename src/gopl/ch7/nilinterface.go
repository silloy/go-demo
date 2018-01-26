package main

import (
	"bytes"
	"io"
)

const debug  = true

func main()  {
	//var buf *bytes.Buffer
	var buf io.Writer  // modify
	if debug {
		buf = new(bytes.Buffer)
	}

	f(buf)

	if debug {

	}
}
func f(out io.Writer) {
	// debug 为false时 out变量是一个包含空指针值的非空接口
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
