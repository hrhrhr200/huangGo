package main

import "fmt"

func printGo() {
	s := "G"

	for i := 0; i < 10; i++ {
		fmt.Println(s)
		s += "G"
	}
}

func main() {
	printGo()
}
