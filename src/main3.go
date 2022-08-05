package main

import "fmt"

func main() {
	//创建一个map必须要用make，否则会是nil
	//格式为:  make(map[key类型]value类型)
	//Java中:   Map<String,Integer> myMap = new HashMap<>();
	myMap := make(map[string]int)
	//往Go中的map赋值添加元素用 【 map变量名称[key] = value 】 的方式
	//区别于Java中的: myMap.put("li_age",20);
	myMap["li_age"] = 20
	myMap["hong_age"] = 30
	//打印一下map
	fmt.Println(myMap)
	//不存在的值
	fmt.Println(myMap["no"])

	//当我们的key取不到对应的值，而value的类型是一个int类型，我们如何判断这个0是实际值还是默认值呢
	//此时我们需要同时取两个值
	//通过map的key取出两个值，第二个参数为bool类型，false为该值不存在，true为成功取到值
	value, existsValue := myMap["no"]
	if !existsValue {
		fmt.Println("此值不存在")
	} else {
		fmt.Printf("value = %d", value)
	}
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	delete(myMap, "li_age")
	for key, value := range myMap {
		fmt.Println(key, value)
	}
	//类似于Java中的Map<String,HashMap<String,Object>>
	var map2 = make(map[string]map[string]interface{})
	map2["huangrui"] = make(map[string]interface{}, 3)
	map2["huangrui"]["a"] = 5
	map2["huangrui"]["b"] = 6
	//map2["huangrui"]["c"] = 7
	//map2["huangrui"]["d"] = 8
	fmt.Println(map2)
}
