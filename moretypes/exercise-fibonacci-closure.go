/*
练习：斐波纳契闭包
让我们用函数做些好玩的事情。

实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。
*/

package main

import "fmt"

func fibonacci() func() int {
	num := []int{0, 1}
	id := 0
	return func() int {
		num = append(num, num[id]+num[id+1])
		id++
		return num[id+1]
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}
