package main

import (
	"flag"
	"io"
	"os"
	"fmt"
	"strings"
)

var (
	n = flag.Bool("n", false, "omit trailing newling")
	s = flag.String("s", " ", "seporator")
)

var out io.Writer = os.Stdout

func main()  {
	flag.Parse()
	if err := echo(!*n, *s, flag.Args()); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func echo(newline bool, seq string, args []string) error {
	fmt.Fprint(out, strings.Join(args, seq))
	if newline {
		fmt.Fprintln(out)
	}
	return nil
}
