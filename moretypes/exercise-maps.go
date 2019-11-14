/*
实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。

你会发现 strings.Fields 很有帮助。
*/

package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	words := strings.Fields(s)
	for _, v := range words {
		m[v] = len(v)
	}
	fmt.Println(m)
	return m
}

func main() {
	WordCount("Prepared for any insane adventure life throws our way")
}
