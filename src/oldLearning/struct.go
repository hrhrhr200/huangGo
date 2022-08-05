package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式的赋予
	v3 = Vertex{}      //X:0 Y:0
	p  = &Vertex{1, 2} //创建一个*Vertex类型的结构体

)

func main() {
	fmt.Println(v1, v2, v3, *p)

	fmt.Println(Vertex{1, 2})

	//结构体字段
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	//结构体指针
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
