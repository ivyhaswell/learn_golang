/*
for（续）
初始化语句和后置语句是可选的。
*/

package main

import "fmt"

func main() {
	sum := 0
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}
