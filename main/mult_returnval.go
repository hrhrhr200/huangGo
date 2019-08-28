package main

import "fmt"

func cula(x, y int) (int, int, int) {
	return x + y, x - y, x * y
}

func cula2(x, y int) (a, b, c int) {
	a = x + y
	b = x - y
	c = x * y
	return
}

func main() {
	a, b, c := cula(3, 4)
	fmt.Println(a, b, c)
	x, y, z := cula2(3, 4)
	fmt.Println(x, y, z)
}
