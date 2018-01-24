package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func add1(r rune) rune { return r + 1 }

//func main() {
//	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
//	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
//	fmt.Println(strings.Map(func(r rune) rune { return r + 1 },
//	"Admix"))    // "Benjy"
//	outlines(os.Args[1])
//}

func outlines(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

// forEachNode针对每个结点x,都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前,pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
