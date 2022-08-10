package main

import (
	"fmt"
	"reflect"
)

func main() {

	/*var num = 50
	testRef(&num)
	fmt.Println(num)*/
	var s = Student{Age: 1, Name: "huangrui"}
	testRef(s)
}

func testRef(i interface{}) {

	t := reflect.TypeOf(i)
	fmt.Printf("传入参数类型为%v", t)
	fmt.Println()
	fmt.Printf("传入参数类型为%v  ", t.Kind())
	fmt.Println()
	v := reflect.ValueOf(i)
	/*fmt.Println(v)
	v.Elem().SetInt(60)
	fmt.Println(v)*/
	n := v.NumField()
	for i := 0; i < n; i++ {
		f := v.Field(i)
		fmt.Printf("该实体类的属性值为%v\n", f)
	}
	mn := v.NumMethod()
	for i := 0; i < mn; i++ {
		m := v.Method(i)
		fmt.Printf("该实体类的方法值为%v\n", m)
		//m.Call()
	}
	v.Method(0).Call(nil)
	ia := make([]reflect.Value, 1)
	if s, flag := i.(Student); flag {
		ia = append(ia, reflect.ValueOf(s))
		v.Method(1).Call(nil)
	}

}

type Student struct {
	Age  int
	Name string
}

func (s Student) Abc() {
	fmt.Println("执行啊方法！！")
}

func (s Student) Bbc() {
	fmt.Println(s.Age)
}

func (s Student) Cbc(a int) {
	s.Age = a
	fmt.Println("c" + string(s.Age))
}
