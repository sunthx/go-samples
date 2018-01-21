package main

import "fmt"

func main() {
	massForPlanet := make(map[string]float64)
	massForPlanet["sunth"] = 0.06
	massForPlanet["sunth1"] = 0.82
	massForPlanet["sunth2"] = 0.83
	massForPlanet["sunth3"] = 0.84

	fmt.Println(massForPlanet)

	triangle := make(map[*Point]string, 3)
	triangle[&Point{89,34,33}] = "1"
	triangle[&Point{11,22,33}]="2"
	triangle[&Point{22,33,44}] = "3"
	fmt.Println(triangle)
	fmt.Println(triangle)

	nameForPoint := make(map[Point]string)
	nameForPoint[Point{54,91,33}] = "1"
	nameForPoint[Point{11,22,33}] = "2"
	nameForPoint[Point{22,33,44}] = "3"
	fmt.Println(nameForPoint)

	map1 := map[string]int{"value1":1,"value2":2,"value3": 3}
	for name,value := range map1{
		fmt.Printf("%-10s %8d\n",name,value)
	}
}

type Point struct {
	x,y,z int
}

func (point Point) String() string  {
	return fmt.Sprintf("(%d%d%d)",point.x,point.y,point.z)
}

