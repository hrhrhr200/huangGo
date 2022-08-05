package main

import "fmt"

func main() {
	i, j := 42, 2701
	p := &i         //指向i
	fmt.Println(*p) //打印p指针指向的底层指
	*p = 21         //通过指针设置i的值
	fmt.Println(i)  //查看i的值

	p = &j         //指向j
	*p = *p / 37   //通过指针对j进行除法运算
	fmt.Println(j) //查看j的值
}
