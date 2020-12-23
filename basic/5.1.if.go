package main

import "fmt"

func main() {
	a := 0

	// 变量 v 的作用域在 if 语句内
	if v := 1; v > 2 {
		a = v + 10
	} else if v > 1 {
		a = v + 100
	} else {
		a = v + 1000
	}

	fmt.Println(a)
}
