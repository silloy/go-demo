package main

import (
	"runtime"
	"fmt"
	"math"
	"bufio"
	"os"
	"strings"
)

type polar struct {
	radius float64
	θ      float64
}
type cartesian struct {
	x float64
	y float64
}

var promt = "Enter a radius and an angle (in degress), eg. 12.5 90," + "os %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		promt = fmt.Sprint(promt, "Ctrl+z, Enter")
	} else {
		promt = fmt.Sprint(promt, "Ctrl+D")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

const result = "Polar radius = %.02f, θ = %.02f° → Cartesian x = %.02f y = %.02f\n"

func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(promt)
	for {
		fmt.Println("Redius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var radius, θ float64
		if _, err := fmt.Sscan(line, "%f %f", &radius, &θ); err != nil {
			fmt.Println(os.Stderr, "invalid input")
			continue
		}
		questions <- polar{radius, θ}
		coord := <-answers
		fmt.Printf(result, radius, θ, coord.x, coord.y)
	}
	fmt.Println()
}

func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarcard := <-questions
			θ := polarcard.θ * math.Pi / 180.0
			x := polarcard.radius * math.Cos(θ)
			y := polarcard.radius * math.Sin(θ)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}


type BitFlag int

const (
	Active BitFlag = 1 << iota
	Send
	Receive
)

//flag := Active | Send

func (flag BitFlag) String() string {
	var flags  []string
	if flag & Active {
		flags = append(flags, "Active")
	}
	if flag & Send {
		flags = append(flags, "Send")
	}
	if flag & Receive {
		flags = append(flags, "Receive")
	}
	if len(flags) {
		return fmt.Sprintf("%d(%s)", int(flag), strings.Join(flags, "|"))
	}
	return "0()"
}
