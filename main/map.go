package main

import "fmt"

type ver struct {
	lat, long float64
}

var m = map[string]ver{
	"Bell Labs": ver{
		40.68433, -74.39967,
	},
	"Google": ver{
		37.42202, -122.08408,
	},
}

func main() {
	/*var m = make(map[string]ver)

	m["aaa"] = ver{
		lat:  222,
		long: 111,
	}*/

	fmt.Println(m)
}
