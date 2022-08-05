package main

import (
	"fmt"
	"gocode/testproject01/demo3/extm2"
)

func main() {
	var c = &extm2.Cat{&extm2.Animal{Age: 12, Heavy: 24.6}}
	fmt.Println(c)
	c.Age = 50
	c.Heavy = 58.4
	c.Hello()
}
