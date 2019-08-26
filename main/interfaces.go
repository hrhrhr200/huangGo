package main

import (
	"fmt"
	"math"
)

type huang interface {
	abs() float64
}

type huang2 float64

func (f huang2) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type huang3 struct {
	X, Y float64
}

func (v *huang3) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var h huang
	//f := huang2(-math.Sqrt2)
	v := huang3{3, 4}
	//h = f
	h = &v
	//h = v
	fmt.Println(h.abs())

}
