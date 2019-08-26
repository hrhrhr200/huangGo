package main

import "fmt"

type bb struct {
	x, y float64
}

func (v *bb) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func scale2(v *bb, f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := bb{3, 4}
	v.scale(2)
	scale2(&v, 2)
	fmt.Println(v)

	p := &bb{4, 5}
	scale2(p, 2)
	fmt.Println(*p)
}
