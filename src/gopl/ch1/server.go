package main

import (
	"net/http"
	"log"
	"fmt"
	"sync"
)

var mu sync.Mutex
var count int

// 字符串转换为数字可以调用strconv.Atoi函数
func main() {
	http.HandleFunc("/", handler);
	http.HandleFunc("/li", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	});
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8800", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)

	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Patth = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

type Point struct {
	X, Y int
}

