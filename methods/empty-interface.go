/*
空接口
指定了零个方法的接口值被称为 *空接口：*

interface{}
空接口可保存任何类型的值。（因为每个类型都至少实现了零个方法。）

空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{} 的任意数量的参数。
*/

package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i interface{}
	describe(i)

	i = 13
	describe(i)

	i = "Hello"
	describe(i)
}
