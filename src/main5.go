package main

import (
	"fmt"
	"sync"
)

func main() {

	myChan := make(chan int, 1)
	//select 的用法是，从上到下依次判断case 是否可执行，如果可执行，则执行完毕跳出select,如果不能执行，尝试下一个执行
	//这里的可执行是指的不阻塞，也就是说，select从上到下开始挑选一个不阻塞的case执行，执行完毕后跳出，
	//如果所有case都阻塞，则执行default
	//如下输出结果，i=奇数的时候走case   myChan<-i:，把奇数放入mychan
	//走偶数的时候因为myChan中有数据了，则把上一个奇数打印出来。
	//所以结果是 1  3  5  7  ...
	for i := 1; i < 100; i++ {
		select {
		case data := <-myChan:
			fmt.Println(data)
		case myChan <- i:
		default:
			fmt.Println("default!!!")
		}
	}

	wait.Add(5)
	go myAdd()
	go myAdd()
	go myAdd()
	go myAdd()
	go myAdd()
	wait.Wait()
	fmt.Printf("num=%d\n", num)
}

var num int
var wait sync.WaitGroup

var lock2 sync.Mutex

func myAdd() {
	defer wait.Done()
	for i := 0; i < 10000; i++ {
		lock2.Lock()
		num += 1
		lock2.Unlock()
	}
}
