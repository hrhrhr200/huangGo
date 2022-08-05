package main

import "fmt"

func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

func cb() (i int) {
	defer func() { i++ }()
	return 1
}

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
	b()
	fmt.Println(cb())

	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
