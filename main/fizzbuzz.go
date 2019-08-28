package main

import "fmt"

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0:
			fmt.Println("fizz")
			fallthrough
		case i%5 == 0:
			fmt.Println("buzz")
			fallthrough
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizzbuzz")
			fallthrough
		default:
			fmt.Println(i)
		}
	}
}

func printrec() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 19; j++ {
			fmt.Print("*")
		}
		fmt.Println("*")
	}
}

func main() {
	//fizzbuzz()
	printrec()
}
