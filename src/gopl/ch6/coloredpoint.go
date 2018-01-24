package main

import (
	"image/color"
	"fmt"
	"sync"
)



//type ColoredPoint struct {
//	Point
//	Color color.RGBA
//}
//
//func main()  {
//	var cp ColoredPoint
//	cp.X = 1
//	fmt.Println(cp.Point.X)
//	cp.Y = 2
//	fmt.Println(cp.Y)
//
//	var p = ColoredPoint{Point{1, 1}, color.RGBA{255, 0, 0, 255}}
//	var q = ColoredPoint{Point{5, 4}, color.RGBA{0, 0, 255, 255}}
//	fmt.Println(p.Distance(q.Point)) // "5"
//	p.ScaleBy(2)
//	q.ScaleBy(2)
//	fmt.Println(p.Distance(q.Point)) // "10"
//}



type ColoredPoint struct {
	*Point
	Color color.RGBA
}

func main()  {
	p := ColoredPoint{&Point{1, 1}, color.RGBA{255, 0, 0, 255}}
	q := ColoredPoint{&Point{5, 4}, color.RGBA{0, 0, 255, 255}}
	fmt.Println(p.Distance(*q.Point)) // "5"
	q.Point = p.Point                 // p and q now share the same Point
	p.ScaleBy(2)
	fmt.Println(*p.Point, *q.Point) // "{2 2} {2 2}"
}

var (
	mu sync.Mutex // guards mapping
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}


// 使用匿名字段
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}


func Lookups(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}