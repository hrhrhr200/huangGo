package main

import (
	"gocode/testproject01/demo4/asset"
)

func main() {
	var s = asset.Chinese{}
	Ger(s)
	s.Sayhello()

	var a = asset.America{}
	Ger(a)

}

func Ger(p asset.Person) {
	p.Sayhello()
	/*if ch, flag := p.(asset.Chinese); flag {
		ch.Dance()
	} else {
		fmt.Println("不是中国人!")
	}*/
	switch p.(type) {
	case asset.Chinese:
		var ch = p.(asset.Chinese)
		ch.Dance()
	case asset.America:
		var a = p.(asset.America)
		a.Play()
	}

}
