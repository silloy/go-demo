package main

import (
	"time"
	"io"
)

func main() {
	//var w io.Writer
	//w = os.Stdout           // OK: *os.File has Write method
	//w = new(bytes.Buffer)   // OK: *bytes.Buffer has Write method
	//w = time.Second         // compile error: time.Duration lacks Write method

	//var rwc io.ReadWriteCloser
	//rwc = os.Stdout         // OK: *os.File has Read, Write, Close methods
	//rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method

	//w = rwc                 // OK: io.ReadWriteCloser has Write method
	//rwc = w                 // compile error: io.Writer lacks Close method

}

type Artiface interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}

type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
	Resolution() (x, y int)
}

type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}


