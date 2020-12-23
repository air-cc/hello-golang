package main

import "fmt"

/**
	延迟函数
	defer 语句会将函数推迟到外层函数返回之后执行。
	推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用
	后进先出
	执行结果:
	`
	1.hello 1
	2.world 0
	3.bye 0
	`
**/
func main() {
	i := 0

	defer fmt.Println("3. bye", i)
	defer fmt.Println("2. world", i)

	i++

	fmt.Println("1. hello", i)

	var j int
	for ; j < 3; j++ {
		defer fmt.Println(j)
	}

	fmt.Println("j", j)
}
