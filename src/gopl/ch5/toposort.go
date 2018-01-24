package main

import (
	"fmt"
	"sort"
	"os"
)

/**
 给定一些计算机课程，每个课程都有前置课程，只有完成了前置课程才可以开始当前课程的学习；
我们的目标是选择出一组课程，这组课程必须确保按顺序学习时，能全部被完成
 */
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

// 使用go语句（第八章）或者defer语句（5.8节）会经常遇到此类问题。
// 这不是go或defer本身导致的，而是因为它们都会等待循环结束后，再执行函数值
//func createAndRemove()  {
//	var rmdirs []func()
//	for _, d := range tempDirs() {
//		dir := d // NOTE: necessary!
//		os.MkdirAll(dir, 0755) // creates parent directories too
//		rmdirs = append(rmdirs, func() {
//			os.RemoveAll(dir)
//		})
//	}
//	// ...do some work…
//	for _, rmdir := range rmdirs {
//		rmdir() // clean up
//	}
//}
