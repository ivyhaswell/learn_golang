/*

练习：Reader
实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流。
*/

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
func (r MyReader) Read([]byte) (int, error) {
	return int('A'), nil
}

func main() {
	reader.Validate(MyReader{})
}
