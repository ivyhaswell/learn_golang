/*
映射的文法（续）
若顶级类型只是一个类型名，你可以在文法的元素中省略它。
*/

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {
		12.123, 42.234,
	},
	"Google": {
		542.123, 34.098,
	},
}

func main() {
	fmt.Println(m)
}
