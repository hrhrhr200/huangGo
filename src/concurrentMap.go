package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var myConcurrentMap = sync.Map{}

var myRangeMap = sync.Map{}

func init() {
	fmt.Println("初始化函数！！！")
}

func main() {
	//存储数据
	myConcurrentMap.Store(1, "huangrui")
	//取数据
	name, ok := myConcurrentMap.Load(1)
	if !ok {
		fmt.Println("不存在")
		return
	}
	fmt.Println(name)
	//该key有值,则ok为true,返回它原来存在的值，不做任何操作；该key无值，则执行添加操作，ok为false,返回新添加的值
	name2, ok2 := myConcurrentMap.LoadOrStore(1, "siqi")
	fmt.Println(name2, ok2)
	name3, ok3 := myConcurrentMap.LoadOrStore(2, "xiao_hong")
	//因为key=2不存在，所以打印是   xiao_hong false
	fmt.Println(name3, ok3)
	//删除值
	myConcurrentMap.Delete(1)

	rangeFunc()

	//原子类，也是cas
	var num10 int64 = 20
	atomic.AddInt64(&num10, 1)
	fmt.Println(num10)
}

//遍历
func rangeFunc() {
	myRangeMap.Store(1, "xiao_ming")
	myRangeMap.Store(2, "xiao_li")
	myRangeMap.Store(3, "xiao_ke")
	myRangeMap.Store(4, "xiao_lei")

	myRangeMap.Range(func(k, v interface{}) bool {
		fmt.Println("data_key_value = :", k, v)
		//return true代表继续遍历下一个，return false代表结束遍历操作
		return false
	})

}
