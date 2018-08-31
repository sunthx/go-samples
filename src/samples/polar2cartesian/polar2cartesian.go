package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

var prompt = "Enter a radius and an angle (in degrees),e.g.,12.5 90,or %s to quit."

func init()  {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt,"Ctrl+Z,Enter")
	} else {
		prompt = fmt.Sprintf(prompt,"Ctrl + D")
	}
}

func main()  {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions,answers)
}

const result  = "Polar radius = %.02f θ=%.02f° → Cartesian x=%.02f y=.02f\n"
func interact(question chan polar,answer chan cartesian) {
	reader:=bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for{
		fmt.Println("Radius and angle:")
		line,err := reader.ReadString('\n')
		if err != nil{
			break
		}
		var radius,θ float64
		fmt.Println(line)
		if _,err := fmt.Sscanf(line,"%f %f",&radius,&θ); err!=nil{
			fmt.Println(os.Stderr, "invalid input")
			continue
		}

		question <- polar{θ, radius}
		coord := <-answer
		fmt.Printf(result,radius,θ,coord.x,coord.y)
	}
	fmt.Println()
}

func createSolver(questions chan polar) chan cartesian{
	answers := make(chan cartesian)
	go func() {
		for{
			polarCoord := <- questions
			θ := polarCoord.θ * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(θ)
			y := polarCoord.radius * math.Sin(θ)
			answers <- cartesian{x,y}
		}
	}()
	return answers
}

type polar struct {
	θ 			float64
	radius 		float64
}

type cartesian struct {
	x float64
	y float64
}
