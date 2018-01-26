package main

import (
	"time"
	"fmt"
)

type Buffer struct {
	buf     []byte
	initial [64]byte
}

func (b *Buffer) Grow(n int)  {
	if b.buf == nil {
		b.buf = b.initial[:0]
	}

	if len(b.buf) + n > cap(b.buf) {
		buf := make([]byte, b.Len(), 2 * cap(b.buf) + n)
		copy(buf, b.buf)
		b.buf = buf
	}
}


func (b *Buffer) Len() int {
	return 0
}

func main()  {
	const day = 24 * time.Hour
	fmt.Println(day.Seconds()) // "86400"
}