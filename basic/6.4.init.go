// 每个包中的 init 函数在第一次导入时会被自动调用

package main

import "fmt"

var i = 1

func init() {
	i = 2
}

func main() {
	defer func() {
		fmt.Println("i:", i)
	}()

	// i = 3
}
