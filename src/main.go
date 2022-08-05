package main

import "fmt"

func main() {
	/*
	 单分支结构语法格式如下:
	         if 条件判断 {
	            //代码块
	         }
	*/
	/*var num int
	fmt.Printf("请输入数字")
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	if num > 10 {
		fmt.Println("你输入的数字大于10")
	} else if num == 10 {
		fmt.Println("你输入的数字等10于10")
	} else {
		fmt.Println("你输入的数字小于10")
	}*/

	fmt.Println("switch语法！！！！！！！！")

	/*
	  switch 分支语法
	*/

	name := "aaa"

	switch name {
	case "huangrui":
		fmt.Println("hello  " + name)
	case "siqi":
		fmt.Println("hello22222  " + name)
	default:
		fmt.Println("no hello")
	}

	name2 := "bbb"

	switch name2 {
	case "huangrui", "bbb":
		fmt.Println("hello" + name2)
	default:
		fmt.Println("aaaaaaa")
	}

	name3 := 5
	switch name3 {
	case 1:
		fmt.Println("aaaaa")
	case 2:
		fmt.Println("bbbbb")
	case 5:
		fmt.Println("ccccc")
		fallthrough
	case 6:
		fmt.Println("ddddddd")
	default:
		fmt.Println("cccccdfrrwegre")
	}

	fmt.Println("for语法！！！！！！！！")

	/*
		for 语法
	*/
	/*for i := 1; i < 10; i++ {
			fmt.Println(i)
		}
		i := 0
		for i < 3 {
			i++
		}
		for i == 3 {
			fmt.Println(i)
			break
		}

		var data [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		for i, num := range data {
			fmt.Println(i, num)
		}
		if i > 10 {
			goto huang
		}
		fmt.Println("不执行我了嘛")

	huang:
		j := 5
		for ; j < 10; j++ {
			fmt.Println(j + 5)
		}
	*/

	/*
		一维数组
	*/
	thisArr := [2]int{1, 2}
	fmt.Println(thisArr)

	/*
	   二维数组
	   第二个中括号不可以...
	*/
	thisArr2 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	thisArr3 := [...][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println(thisArr2)
	fmt.Println(thisArr3)

	/*
		指针概念
	*/
	name4 := "huangrui"
	var namePoint *string
	namePoint = &name4
	fmt.Println("name 变量的地址是:", &name4)
	fmt.Println("namePoint 变量存储的地址是", namePoint)
	fmt.Println("name的值是", *namePoint)

	var num = new(int)
	fmt.Println(num)
	*num = 1
	fmt.Println(*num)

	//make只能为map,channel,slice申请分配内存，只有这三种，没有第四种
	//所有通过make创建的这三种类型都是引用类型，传递参数时虽然是引用值传递，
	//但是对方法内引用变量参数的修改可以影响到外部的引用变量
	//1.通过make创建map对象 如下代码类似于Java中 Map<String,Integer> myMap = new HashMap<>();
	//在这里make就是申请分配map的内存，和java中创建map的new一样
	myMap := make(map[string]int)
	myMap["huangrui"] = 50
	myMap["siqi"] = 60
	fmt.Println(myMap["huangrui"])

	//2.通过make创建channel,make函数内可以有一个参数，也可以有两个参数，有两个参数时第二个参数
	//为通道的缓存队列的长度
	//2.1) 只有一个参数，通道的缓存队列长度此时为0，也就是无缓存。
	//创建一个传输int类型数据的通道
	myChan := make(chan int)
	fmt.Println(myChan)
	//2.2) 有两个参数，第二个参数2代表此时代表缓存队列的长度为2
	//创建一个传输int类型数据的通道,缓存为2
	mychan2 := make(chan int, 2)
	fmt.Println(mychan2)
	//此处暂时不做通道缓存队列数多少有何区别的讲解
	//我们如果是想创建一个空的slice,则用make创建切片
	//如下形式 make(int[],num1,num2)
	//num1 = 切片的长度(默认分配内存空间的元素个数)
	//num2 = 切片的容量(解释：底层数组的长度/切片的容量，超过底层数组长度append新元素时会创建一个新的底层数组，
	//不超过则会使用原来的底层数组)

	//代表底层数组的长度是4,默认给底层数组的前两个元素分配内存空间
	//切片指向前两个元素的地址，如果append新元素，在元素数小于4时都会
	//在原来的底层数组的最后一个元素新分配空间和赋值，
	//append超过4个元素时，因为原数组大小不可变，也也存储不下了，
	//所以会新创建一个新的底层数组，切片指向新的底层数组
	mySlice := make([]int, 2, 4)
	fmt.Println(mySlice)
	mySlice[0] = 5
	mySlice[1] = 5
	mySlice = append(mySlice, 3, 5, 6, 7)
	fmt.Println(mySlice[4])
	fmt.Println(mySlice)
	print(mySlice)
	mySlice2 := make([]int, 3)
	fmt.Println(mySlice2)

	//函数
	fmt.Println("函数！！！！！！！！！！！")
	var num1 int = 100
	var num2 int = 200
	var result int
	result = max(num1, num2)
	fmt.Println("大数是", result)
	a := "huangrui"
	b := "siqi"
	a, b = swap(a, b)
	fmt.Println("第一个数", a, "第二个数", b)

	c, d := 200, 100
	var data = []string{"11", "22", "33"}
	c, d = manyArgs(c, d, data...)
	fmt.Println(c, d)

	defertest()
	_, num3, _ := testFun(1, 2, "huangrui")
	fmt.Println(num3)

	/*
			close          关闭channel

		len            求长度

		make	      创建slice,map,chan对象

		append	      追加元素到切片(slice)中

		panic         抛出异常，终止程序

		recover       尝试恢复异常，必须写在defer相关的代码块中
	*/
	fmt.Println("常用关键字!!!!")
	arr := [5]int{1, 2, 3, 4, 5}
	str := "huangrui"
	//获取字符串长度
	fmt.Println(len(str))
	//获取数组长度
	fmt.Println(len(arr))
	//获取切片长度
	fmt.Println(len(arr[:3]))

	//创建channel
	intChan := make(chan int, 5)
	//创建map
	map1 := make(map[string]interface{})
	//创建切片
	mySlice3 := make([]int, 5, 10)
	fmt.Println(intChan)
	fmt.Println(map1)
	fmt.Println(mySlice3)

	close(intChan)
	//为切片添加元素
	array2 := append(mySlice3[:], 6, 6, 6, 6, 6, 6)
	fmt.Println(array2)

	num4 := new(int)
	fmt.Println(num4)

	//捕获异常和尝试恢复程序
	func1()
	func2()
	func3()
}

func max(num1, num2 int) int {
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func swap(a, b string) (string, string) {
	return b, a
}

/*
  变长参数
*/
func manyArgs(a int, b int, str ...string) (int, int) {
	for i, s := range str {
		fmt.Println(i, s)
	}

	return b, a
}

/*
	延迟执行 类型于java中的finally
*/
func defer1() {
	fmt.Println("print defer1")
}

func defer2() {
	fmt.Println("print defer2")
}

func defertest() {
	defer defer1()
	defer defer2()
	fmt.Println("defetteststaret111!@!")
}

func testFun(num1, num2 int, str string) (n1 int, n2 int, s1 string) {
	n1 = num1
	n2 = num2
	s1 = str
	return
}
func testFun2(num1, num2 int, str string) (int, int, string) {
	return num1, num2, str
}
func func1() {
	fmt.Println("1")
}

func func2() {
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放资源")
	}()
	//num := 1 / 0
	//fmt.Println(num)
	panic("抛出异常")
	fmt.Println("2")
}

func func3() {
	fmt.Println("3")
}
