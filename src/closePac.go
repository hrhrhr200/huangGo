package main

import "fmt"

var myfunc = func(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(myfunc(1, 2))

	var f = func() {
		fmt.Println("内部函数")
	}
	f()

	func() {
		fmt.Println("立即执行的函数")
	}()

	var f1 = closeP(100)
	f1(100)
	f1(100)
	f1(100)
}

/*
	闭包
*/
func closeP(x int) func(int) {
	return func(y int) {
		x += y
		fmt.Printf("%d数值为\n", x)
	}
}

/*
做下闭包的总结，如何实现一个闭包：

1.定义一个A函数，此函数返回一个匿名函数。（定义一个返回匿名函数的A函数）

2.把在A函数的b参数或A函数代码块中的b变量，放入匿名函数中，进行操作。

3.这样我们调用A函数返回一个函数，这个函数不断的调用就可以一直使用之前b参数，b变量，并且b值不会刷新，有点像在匿名函数外部自定义了一个b的成员变量（成员变量取自Java中类的相关概念）
*/
