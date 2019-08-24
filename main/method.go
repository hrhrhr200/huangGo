package main

import (
	"fmt"
	"math"
)

type number struct {
	x, y float64
}

type MyFloat float64

func (f MyFloat) acs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (a number) abs() float64 {
	return math.Sqrt(a.x*a.x + a.y*a.y)
}

func main() {
	v := number{3, 5}
	fmt.Println(v.abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.acs())
}
