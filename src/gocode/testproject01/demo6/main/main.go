package main

import (
	"fmt"
)

func main() {

	/*go func() {
		for i := 1; i < 10; i++ {
			fmt.Println("gorount hello golant")
			time.Sleep(1 * time.Second)
		}
	}()
	for i := 1; i < 10; i++ {
		fmt.Println("hello golang!!!")
		//time.Sleep(1 * time.Second)
	}*/

	var intChan = make(chan int, 100)

	for i := 1; i < 101; i++ {
		intChan <- i
	}
	//循环读取管道的数据前需要关闭管道，不然会出现死锁现象
	close(intChan)

	for v := range intChan {
		fmt.Println(v)
	}

}
