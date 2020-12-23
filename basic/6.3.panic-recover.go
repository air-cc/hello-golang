// 异常处理

package main

import "fmt"

func throwPanic(f func()) (ret bool) {
	defer func() {
		if result := recover(); result != nil {
			fmt.Println("recover a panic:", result)
			ret = true
		}
	}()

	if f == nil {
		panic("panic happened")
	}

	return
}

func main() {
	var f func()
	// f = func() {}

	ret := throwPanic(f)
	fmt.Println("ret", ret)
}
