package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)   // 函数的并发执行方式
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	for range os.Args[1:] {
		fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	}
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	//if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
	//	url = "https://" + url
	//}
	resp, err := http.Get(url)
	//fmt.Println(resp.Status, ",", url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
