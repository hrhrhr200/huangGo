package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y = 4, 5
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	i := 42
	a := float64(i)
	u := uint(a)
	fmt.Println(u)
}
