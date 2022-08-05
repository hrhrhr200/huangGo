package main

import (
	"fmt"
	"time"
)

func main() {
	//获取当前时间
	now := time.Now()
	//获取当前年份
	year := now.Year()
	//获取当前月份
	month := now.Month()
	//获取当前 日期
	day := now.Day()
	//获取当前小时
	hour := now.Hour()
	//获取当前分钟
	min := now.Minute()
	//获取当前秒
	second := now.Second()

	//获取当前时间戳，和其它编程语言一样，自1970年算起
	timestamp := now.Unix()
	//纳秒时间戳
	ntimestamp := now.UnixNano()

	fmt.Println("year=", year)
	fmt.Println("month=", month)
	fmt.Println("day=", day)
	fmt.Println("hour=", hour)
	fmt.Println("min=", min)
	fmt.Println("second=", second)
	fmt.Println("timestamp=", timestamp)
	fmt.Println("ntimestamp=", ntimestamp)

	//2006-01-02 15:04:05这个数值是一个标准写死的，只要改格式符号即可
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006/01/02")) //年月日
	fmt.Println(now.Format("15:04:05"))   //时分秒
}
