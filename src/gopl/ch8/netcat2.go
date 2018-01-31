package main

import (
	"net"
	"log"
	"os"
	"io"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopyy(os.Stdout, conn)
	mustCopyy(conn, os.Stdin)
}





func mustCopyy(dest io.Writer, src io.Reader) {
	if _, err := io.Copy(dest, src); err != nil {
		log.Fatal(err)
	}
}