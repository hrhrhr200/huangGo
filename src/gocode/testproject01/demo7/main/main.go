package main

import (
	"fmt"
	"sync"
)

func main() {

	//读写双用
	//var mychan chan int

	//只写
	/*var mychan chan<- int
	mychan = make(chan int, 3)
	mychan <- 20*/
	//num := <-mychan
	//fmt.Println(num)

	//只读
	//var mychan2 <-chan int

	waitGroup.Add(2)

	var two = func() {
		defer waitGroup.Done()
		if money > 200 {
			sub(200)
			fmt.Println("减去了200元")
		}
	}
	go two()

	var seven = func() {
		defer waitGroup.Done()
		if money > 900 {
			sub(900)
			fmt.Println("减去了900元")
		}
	}

	go seven()

	waitGroup.Wait()
	fmt.Println(getMoney())

}

var lock sync.Mutex

var money = 1000

var waitGroup sync.WaitGroup

func sub(subNum int) {
	lock.Lock()
	money -= subNum
	lock.Unlock()
}

func getMoney() int {
	return money
}
