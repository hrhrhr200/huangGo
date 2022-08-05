package main

import (
	"fmt"
	"sync"
	"time"
)

//金额
var money2 = 1000

//读写锁
var rwLock sync.RWMutex

var wait2 sync.WaitGroup

//添加金额
func addM(sub int) {
	rwLock.Lock()
	time.Sleep(3 * time.Second)
	money2 -= sub
	rwLock.Unlock()
}

func getM() int {
	rwLock.RLock()
	result := money2
	rwLock.RUnlock()
	return result
}

func main() {

	wait2.Add(2)

	go func() {
		defer wait2.Done()
		if getM() > 200 {
			addM(200)
			fmt.Printf("200元扣款成功，剩下：%d元\n", getM())
		} else {
			fmt.Println("余额不足200!!")
		}
	}()

	var nine = func() {
		defer wait2.Done()
		if getM() > 900 {
			addM(900)
			fmt.Printf("900元扣款成功，剩下：%d元\n", getM())
		} else {
			fmt.Println("余额不足900!")
		}
	}
	go nine()

	wait2.Wait()
	fmt.Println(getM())

}
