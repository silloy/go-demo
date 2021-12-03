package tests

import (
	"fmt"
	. "src/go-web-test/method"
	"testing"
)

func TestMethod(t *testing.T) {
	boxes := BoxList {
		NewBox(WriteWidth(4), WriteHeight(4), WriteDepth(4), WriteColor(RED)),
		NewBox(WriteWidth(10), WriteHeight(10), WriteDepth(1), WriteColor(RED)),
		NewBox(WriteWidth(1), WriteHeight(1), WriteDepth(20), WriteColor(YELLOW)),
		NewBox(WriteWidth(10), WriteHeight(10), WriteDepth(1), WriteColor(RED)),
		NewBox(WriteWidth(1), WriteHeight(1), WriteDepth(20), WriteColor(BLUE)),
		NewBox(WriteWidth(10), WriteHeight(10), WriteDepth(1), WriteColor(RED)),
		NewBox(WriteWidth(10), WriteHeight(30), WriteDepth(1), WriteColor(BLUE)),
		NewBox(WriteWidth(20), WriteHeight(20), WriteDepth(20), WriteColor(RED)),
	}

	fmt.Printf("We have %d boxes in our set\n", len(boxes))
	fmt.Println("The volume of the first one is", boxes[0].Volume(), "cm³")
	fmt.Println("The color of the last one is",boxes[len(boxes)-1].Color().String())
	fmt.Println("The biggest one is", boxes.BiggestColor().String())

	fmt.Println("Let's paint them all black")
	boxes.PaintItBlack()
	fmt.Println("The color of the second one is", boxes[1].Color().String())

	fmt.Println("Obviously, now, the biggest one is", boxes.BiggestColor().String())
}