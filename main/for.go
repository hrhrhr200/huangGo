package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//无限循环
	//如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑
	/*for{

	}*/
}
