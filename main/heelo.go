package main

import "fmt"

func main() {

	fv := func() { fmt.Println("hello,World") }
	fv()
	fmt.Printf("aaa%t", fv)
}
