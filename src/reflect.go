package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*testAllType(1)
	testAllType("go")

	handleType(1)
	handleType(true)*/

	handleType2(1)
	handleType2(true)
}

/*
go中有一个类型 interface,
interface类型相当于Java中的Object类，当以interface作为参数类型时，可以给这个参数传递任意类型的变量。
*/

func testAllType(data interface{}) {
	fmt.Println(data)
}

/*
反射我们需要调用reflect包模块,使用reflect.typeOf()可以获取参数的类型信息对象，
再根据类型信息对象的kind方法，获取到具体类型
*/
func handleType(data interface{}) {

	d := reflect.TypeOf(data)
	fmt.Println(d.Kind())
	fmt.Println(d)
	switch d.Kind() {
	case reflect.Invalid:
		fmt.Println("无效类型")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("整形")
	case reflect.Bool:
		fmt.Println("布尔类型")
	}

}

func handleType2(data interface{}) {

	d := reflect.ValueOf(data)
	fmt.Println(d.Kind())
	fmt.Println(d)
	switch d.Kind() {
	case reflect.Invalid:
		fmt.Println("无效类型")
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var num = d.Int()
		fmt.Println("整形", num)
	case reflect.Bool:
		var boo = d.Bool()
		fmt.Println("布尔类型", boo)
	}

}
