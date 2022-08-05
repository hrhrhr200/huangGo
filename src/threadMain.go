package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*go goroutineTestFunc()
	time.Sleep(10 * time.Second)
	myFunction()*/
	/*myFunction()
	go goroutineTestFunc()*/

	/*go starta()
	go startb()
	time.Sleep(10 * time.Second)*/

	go func() {
		//moneyChan <- 10
		if getMoney() > 200 {
			sMoney(200)
			fmt.Println("減去200！！")
		}
		//<-moneyChan
	}()

	go func() {
		//moneyChan <- 10
		if getMoney() > 900 {
			sMoney(900)
			fmt.Println("減去900！！")
		}
		//<-moneyChan
	}()
	time.Sleep(8 * time.Second)
	fmt.Println(getMoney())
}

//我的方法
func myFunction() {
	fmt.Println("Hello!!!")
}

//并发执行方法
func goroutineTestFunc() {
	fmt.Println("Hello!!! Start Goroutine!!!")
}

//channel发送数据和接受数据用 <-表示,是发送还是接受取决于chan在  <-左边还是右边
var chanStr = make(chan string)

func starta() {
	fmt.Println("發送管道數據")
	chanStr <- "hello!!"
}

func startb() {
	fmt.Println("接受管道數據")
	str := <-chanStr
	fmt.Println(str)
}

var money int = 1000

func sMoney(sub int) {
	lock.Lock()
	time.Sleep(3 * time.Second)
	money -= sub
	fmt.Println(money)
	lock.Unlock()
}

func getMoney() int {
	lock.Lock()
	result := money
	lock.Unlock()
	return result
}

var moneyChan = make(chan int, 1)

var lock sync.Mutex
