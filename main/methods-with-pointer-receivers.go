package main

import (
	"fmt"
	"math"
)

type cc struct {
	x, y float64
}

func (v *cc) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func (v *cc) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := &cc{3, 5}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.abs())
	v.scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.abs())
}
