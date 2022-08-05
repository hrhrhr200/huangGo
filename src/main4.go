package main

import "fmt"

type Boy struct {
	name string
	age  int
}

func mod(b *Boy) {
	//这个是获取调用方法传入的参数的地址值
	fmt.Printf("b的值(之前boy的地址)是%p\n", b)
	//这个是获取本函数中 b这个指针变量的地址
	fmt.Printf("b这个指针自己的地址是=%p\n", &b)
	//打印值\
	b.name = "siqi"
	//这里自动转换使指针可以直接点出来属性
	fmt.Println(b.name, b.age)
}

func main() {
	boy := &Boy{"huangrui", 18}
	fmt.Printf("main函数中的boy地址是:%p\n", boy)
	mod(boy)
	fmt.Println(boy.name, boy.age)
	//注意！！！下面有黑箱操作：
	/* //在&boy并放入Mod传递的过程中实际上做了如下黑箱操作
	   b := new(Boy)   //创建一个名为b的类型为Boy的指针变量
	   b = &boy      //把boy的地址存入b这个指针变量内
	   //接着把b放入func Mod(b *Boy)的参数中，然后，开始执行Mod方法。
	   fmt.Println(b.name,b.age)
	   fmt.Printf("b的地址是:%p\n",&b)
	   fmt.Printf("b的值是:%p\n",b)

	   //输出结果
	   //main函数中的boy地址是:0x10aec0c0
	   //li_ming 20
	   //b的地址是:0x10ae40f8
	   //b的值是:0x10aec0c0
	*/

	/*//以下代码无用，是指为了加深理解new，可以试试输出结果
	  boy2 := new(Boy)
	  fmt.Printf("main函数中new的boy2地址是:%p\n",boy2)
	  boy2.name = "xiaohong"
	  boy2.age = 18
	  Mod(boy2)
	*/

	//加法函数赋值给该函数变量，相当于函数指针指向加法函数
	myPointFun = myAddFun
	fmt.Printf("a+b = %d\n", myPointFun(10, 20))
	//减法函数赋值给该函数变量，相当于函数指针指向减法函数
	myPointFun = mySubFun
	fmt.Printf("a-b = %d\n", myPointFun(100, 50))
}

/*
​ 1.Go的参数传递都是值传递。

​ 2.指针类型的值传递可以改变原来对象的值。

​ 3.make和new从底层原理上创建的所有对象都是指针对象，所以make和new创建出来的slice,map,chan或者其它任何对象都是指针传递，改变值后都可以使原来的对象属性发生变化。
*/

//加法
func myAddFun(x, y int) int {
	return x + y
}

//减法
func mySubFun(x, y int) int {
	return x - y
}

//函数变量(类似于函数指针)
var myPointFun func(x, y int) int
