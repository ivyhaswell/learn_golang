/*
映射的文法
映射的文法与结构体相似，不过必须有键名。
*/

package main

import "fmt"

type VerTex struct {
	Lat, Long float64
}

var m = map[string]VerTex{
	"Bell Labs": VerTex{
		34.019238, 38.90283,
	},
	"Google": VerTex{
		34.093, 3.980,
	},
}

func main() {
	fmt.Println(m)
}
