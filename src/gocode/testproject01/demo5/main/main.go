package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
IO
*/
func main() {

	//file, err := os.Open("/Users/viki/test")
	file, err := os.OpenFile("/Users/viki/test", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开出错，对应错误为:", err)
	}
	//读取的另一种方法
	defer file.Close()
	/*fmt.Printf("文件=%v", file)

	err2 := file.Close()
	if err2 != nil {
		fmt.Println("数据源关闭出错！！！")
	}*/

	//读取文件
	//过时方法
	/*content, _ := ioutil.ReadFile("/Users/viki/test")
	fmt.Printf("文件内容:%v", string(content))*/

	//文件中添加数据
	/*writer := bufio.NewWriter(file)
	for i := 1; i < 5; i++ {
		writer.WriteString("写文件操作写写写" + string(i) + "\n")
	}

	writer.Flush()*/

	reader := bufio.NewReader(file)
	for {
		str, readErr := reader.ReadString('\n')
		if readErr == io.EOF {
			break
		}
		fmt.Println(str)
	}
	fmt.Println("文件输出完毕!")
}
