package main

import (
	"fmt"
	//"math"
)

func main() {
	a := 10
	b := 3
	c := a + b
	fmt.Println("Addition:", c)

	//fmt.Println("Min:", math.Min(float64(a), float64(b)))
	fmt.Println("Max:", max(a, b))

}
