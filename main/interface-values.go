package main

import (
	"fmt"
	"math"
)

type huangi2 interface {
	M()
}

type huangi3 struct {
	s string
}

func (i *huangi3) M() {
	if i == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(i.s)
}

type f float64

func (a f) M() {
	fmt.Println(a)
}

func describe(i huangi2) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i huangi2

	var t *huangi3
	i = t
	describe(i)
	i.M()

	i = f(-math.Sqrt2)
	describe(i)
	i.M()
}
